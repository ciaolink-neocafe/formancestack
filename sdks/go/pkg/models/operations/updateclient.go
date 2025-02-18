// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateClientRequest struct {
	UpdateClientRequest *shared.UpdateClientRequest `request:"mediaType=application/json"`
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type UpdateClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Updated client
	UpdateClientResponse *shared.UpdateClientResponse
}
