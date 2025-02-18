package activities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type ConfirmHoldRequest struct {
	ID string `json:"id"`
}

func (a Activities) ConfirmHold(ctx context.Context, request ConfirmHoldRequest) error {
	response, err := a.client.Wallets.ConfirmHold(
		ctx,
		operations.ConfirmHoldRequest{
			ConfirmHoldRequest: &shared.ConfirmHoldRequest{},
			HoldID:             request.ID,
		},
	)
	if err != nil {
		return err
	}

	if response.WalletsErrorResponse != nil {
		return temporal.NewApplicationError(
			response.WalletsErrorResponse.ErrorMessage,
			string(response.WalletsErrorResponse.ErrorCode),
		)
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}

var ConfirmHoldActivity = Activities{}.ConfirmHold

func ConfirmHold(ctx workflow.Context, id string) error {
	return executeActivity(ctx, ConfirmHoldActivity, nil, ConfirmHoldRequest{
		ID: id,
	})
}
