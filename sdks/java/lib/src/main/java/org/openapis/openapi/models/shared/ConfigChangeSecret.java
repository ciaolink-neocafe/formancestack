/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ConfigChangeSecret {
    @JsonProperty("secret")
    public String secret;

    public ConfigChangeSecret withSecret(String secret) {
        this.secret = secret;
        return this;
    }
    
    public ConfigChangeSecret(@JsonProperty("secret") String secret) {
        this.secret = secret;
  }
}
