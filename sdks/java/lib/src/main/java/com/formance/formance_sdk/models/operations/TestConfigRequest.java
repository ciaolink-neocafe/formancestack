/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class TestConfigRequest {
    /**
     * Config ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;

    public TestConfigRequest withId(String id) {
        this.id = id;
        return this;
    }
    
    public TestConfigRequest(@JsonProperty("id") String id) {
        this.id = id;
  }
}
