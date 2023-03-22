/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class GetInstanceHistoryResponse {
    
    public String contentType;

    public GetInstanceHistoryResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * General error
     */
    
    public org.openapis.openapi.models.shared.Error error;

    public GetInstanceHistoryResponse withError(org.openapis.openapi.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    /**
     * The workflow instance history
     */
    
    public org.openapis.openapi.models.shared.GetWorkflowInstanceHistoryResponse getWorkflowInstanceHistoryResponse;

    public GetInstanceHistoryResponse withGetWorkflowInstanceHistoryResponse(org.openapis.openapi.models.shared.GetWorkflowInstanceHistoryResponse getWorkflowInstanceHistoryResponse) {
        this.getWorkflowInstanceHistoryResponse = getWorkflowInstanceHistoryResponse;
        return this;
    }
    
    
    public Integer statusCode;

    public GetInstanceHistoryResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public GetInstanceHistoryResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public GetInstanceHistoryResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}
