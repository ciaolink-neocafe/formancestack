/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * ListUsersResponse - List of users
 */
public class ListUsersResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public User[] data;

    public ListUsersResponse withData(User[] data) {
        this.data = data;
        return this;
    }
    
    public ListUsersResponse(){}
}
