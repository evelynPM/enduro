package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/artefactual-labs/enduro/internal/nha"
	wferrors "github.com/artefactual-labs/enduro/internal/workflow/errors"
	"github.com/artefactual-labs/enduro/internal/workflow/manager"
)

// Similar to time.RFC3339 with dashes and colons removed.
const rfc3339forFilename = "20060102.150405.999999999"

// UpdateProductionSystemActivity sends a receipt to the production system
// using their filesystem interface using GoAnywhere.
type UpdateProductionSystemActivity struct {
	manager *manager.Manager
}

func NewUpdateProductionSystemActivity(m *manager.Manager) *UpdateProductionSystemActivity {
	return &UpdateProductionSystemActivity{manager: m}
}

type UpdateProductionSystemActivityParams struct {
	StoredAt     time.Time
	PipelineName string
	NameInfo     nha.NameInfo
}

func (a *UpdateProductionSystemActivity) Execute(ctx context.Context, params *UpdateProductionSystemActivityParams) error {
	// We expect tinfo.StoredAt to have the zero value when the ingestion
	// has failed. Here we set a new value as it is a required field.
	if params.StoredAt.IsZero() {
		params.StoredAt = time.Now().UTC()
	}

	receiptPath, err := manager.HookAttrString(a.manager.Hooks, "prod", "receiptPath")
	if err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error looking up receiptPath configuration attribute: %v", err))
	}

	var basename = filepath.Join(receiptPath, fmt.Sprintf("Receipt_%s_%s", params.NameInfo.Identifier, params.StoredAt.Format(rfc3339forFilename)))
	var jsonPath = basename + ".json"
	var mftPath = basename + ".mft"

	// Create and open receipt file.
	file, err := os.OpenFile(jsonPath, os.O_RDWR|os.O_CREATE, os.FileMode(0o644))
	if err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error creating receipt file: %v", err))
	}

	// Write receipt contents.
	if err := a.generateReceipt(params, file); err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error writing receipt file: %v", err))
	}

	// Seek to the beginning of the file.
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error resetting receipt file cursor: %v", err))
	}

	_ = file.Close()

	// Final rename.
	if err := os.Rename(jsonPath, mftPath); err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error renaming receipt (json » mft): %v", err))
	}

	return nil
}

func (a UpdateProductionSystemActivity) generateReceipt(params *UpdateProductionSystemActivityParams, file *os.File) error {
	receipt := prodSystemReceipt{
		Identifier: params.NameInfo.Identifier,
		Type:       params.NameInfo.Type.Lower(),
		Accepted:   true,
		Message:    fmt.Sprintf("Package was processed by Archivematica pipeline %s", params.PipelineName),
		Timestamp:  params.StoredAt,
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(receipt); err != nil {
		return fmt.Errorf("encoding failed: %v", err)
	}

	if err := file.Sync(); err != nil {
		return fmt.Errorf("sync failed: %v", err)
	}

	return nil
}

type prodSystemReceipt struct {
	Identifier string    `json:"identifier"` // Original identifier.
	Type       string    `json:"type"`       // Lowercase. E.g. "dpj", "epj", "other" or "avlxml".
	Accepted   bool      `json:"accepted"`   // Whether we have an error during processing.
	Message    string    `json:"message"`    // E.g. "Package was processed by Archivematica pipeline am" or any other error message.
	Timestamp  time.Time `json:"timestamp"`  // RFC3339, e.g. "2006-01-02T15:04:05Z07:00"
}
