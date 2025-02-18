// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ConfirmHoldRequest struct {
	ConfirmHoldRequest *shared.ConfirmHoldRequest `request:"mediaType=application/json"`
	HoldID             string                     `pathParam:"style=simple,explode=false,name=hold_id"`
}

type ConfirmHoldResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}
