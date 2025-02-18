package wallets

import (
	"fmt"

	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	const (
		metadataFlag = "metadata"
	)
	return fctl.NewCommand("create <name>",
		fctl.WithShortDescription("Create a new wallet"),
		fctl.WithAliases("cr"),
		fctl.WithConfirmFlag(),
		fctl.WithArgs(cobra.ExactArgs(1)),
		fctl.WithStringSliceFlag(metadataFlag, []string{""}, "Metadata to use"),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {
			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return errors.Wrap(err, "retrieving config")
			}

			organizationID, err := fctl.ResolveOrganizationID(cmd, cfg)
			if err != nil {
				return err
			}

			stack, err := fctl.ResolveStack(cmd, cfg, organizationID)
			if err != nil {
				return err
			}

			if !fctl.CheckStackApprobation(cmd, stack, "You are about to create a wallet") {
				return fctl.ErrMissingApproval
			}

			client, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return errors.Wrap(err, "creating stack client")
			}

			metadata, err := fctl.ParseMetadata(fctl.GetStringSlice(cmd, metadataFlag))
			if err != nil {
				return err
			}

			request := shared.CreateWalletRequest{
				Name:     args[0],
				Metadata: metadata,
			}
			response, err := client.Wallets.CreateWallet(cmd.Context(), request)
			if err != nil {
				return errors.Wrap(err, "creating wallet")
			}

			if response.WalletsErrorResponse != nil {
				return fmt.Errorf("%s: %s", response.WalletsErrorResponse.ErrorCode, response.WalletsErrorResponse.ErrorMessage)
			}

			if response.StatusCode >= 300 {
				return fmt.Errorf("unexpected status code: %d", response.StatusCode)
			}

			pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln(
				"Wallet created successfully with ID: %s", response.CreateWalletResponse.Data.ID)
			return nil
		}),
	)
}
