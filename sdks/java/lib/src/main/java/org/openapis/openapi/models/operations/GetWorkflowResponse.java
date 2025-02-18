/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class GetWorkflowResponse {
    
    public String contentType;

    public GetWorkflowResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * General error
     */
    
    public org.openapis.openapi.models.shared.Error error;

    public GetWorkflowResponse withError(org.openapis.openapi.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    /**
     * The workflow
     */
    
    public org.openapis.openapi.models.shared.GetWorkflowResponse getWorkflowResponse;

    public GetWorkflowResponse withGetWorkflowResponse(org.openapis.openapi.models.shared.GetWorkflowResponse getWorkflowResponse) {
        this.getWorkflowResponse = getWorkflowResponse;
        return this;
    }
    
    
    public Integer statusCode;

    public GetWorkflowResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public GetWorkflowResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public GetWorkflowResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}
