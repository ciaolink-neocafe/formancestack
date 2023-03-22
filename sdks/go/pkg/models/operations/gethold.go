// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetHoldRequest struct {
	// The hold ID
	HoldID string `pathParam:"style=simple,explode=false,name=holdID"`
}

type GetHoldResponse struct {
	ContentType string
	// Holds
	GetHoldResponse *shared.GetHoldResponse
	StatusCode      int
	RawResponse     *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}
