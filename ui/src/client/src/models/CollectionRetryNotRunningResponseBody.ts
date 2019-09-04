// tslint:disable
// eslint-disable
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * retry_not_running_response_body result type (default view)
 * @export
 * @interface CollectionRetryNotRunningResponseBody
 */
export interface CollectionRetryNotRunningResponseBody {
    /**
     * Is the error a server-side fault?
     * @type {boolean}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    fault: boolean;
    /**
     * ID is a unique identifier for this particular occurrence of the problem.
     * @type {string}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    id: string;
    /**
     * Message is a human-readable explanation specific to this occurrence of the problem.
     * @type {string}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    message: string;
    /**
     * Name is the name of this class of errors.
     * @type {string}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    name: string;
    /**
     * Is the error temporary?
     * @type {boolean}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    temporary: boolean;
    /**
     * Is the error a timeout?
     * @type {boolean}
     * @memberof CollectionRetryNotRunningResponseBody
     */
    timeout: boolean;
}

export function CollectionRetryNotRunningResponseBodyFromJSON(json: any): CollectionRetryNotRunningResponseBody {
    return CollectionRetryNotRunningResponseBodyFromJSONTyped(json, false);
}

export function CollectionRetryNotRunningResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CollectionRetryNotRunningResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fault': json['fault'],
        'id': json['id'],
        'message': json['message'],
        'name': json['name'],
        'temporary': json['temporary'],
        'timeout': json['timeout'],
    };
}

export function CollectionRetryNotRunningResponseBodyToJSON(value?: CollectionRetryNotRunningResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fault': value.fault,
        'id': value.id,
        'message': value.message,
        'name': value.name,
        'temporary': value.temporary,
        'timeout': value.timeout,
    };
}


