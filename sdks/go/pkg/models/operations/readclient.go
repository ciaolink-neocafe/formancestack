// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ReadClientRequest struct {
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type ReadClientResponse struct {
	ContentType string
	// Retrieved client
	ReadClientResponse *shared.ReadClientResponse
	StatusCode         int
	RawResponse        *http.Response
}
