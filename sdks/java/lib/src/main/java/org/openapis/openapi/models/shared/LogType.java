/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.shared;

import com.fasterxml.jackson.annotation.JsonValue;

public enum LogType {
    NEW_TRANSACTION("NEW_TRANSACTION"),
    SET_METADATA("SET_METADATA"),
    REVERTED_TRANSACTION("REVERTED_TRANSACTION");

    @JsonValue
    public final String value;

    private LogType(String value) {
        this.value = value;
    }
}
