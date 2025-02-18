package wallets

import (
	"fmt"
	"io"

	"github.com/formancehq/fctl/cmd/wallets/internal"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewShowCommand() *cobra.Command {
	return fctl.NewCommand("show",
		fctl.WithShortDescription("Show a wallets"),
		fctl.WithAliases("sh"),
		fctl.WithConfirmFlag(),
		fctl.WithArgs(cobra.ExactArgs(0)),
		internal.WithTargetingWalletByID(),
		internal.WithTargetingWalletByName(),
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

			client, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return errors.Wrap(err, "creating stack client")
			}

			walletID, err := internal.RetrieveWalletID(cmd, client)
			if err != nil {
				return err
			}
			if walletID == "" {
				return errors.New("You need to specify wallet id using --id or --name flags")
			}

			response, err := client.Wallets.GetWallet(cmd.Context(), operations.GetWalletRequest{
				ID: walletID,
			})
			if err != nil {
				return errors.Wrap(err, "getting wallet")
			}

			if response.WalletsErrorResponse != nil {
				return fmt.Errorf("%s: %s", response.WalletsErrorResponse.ErrorCode, response.WalletsErrorResponse.ErrorMessage)
			}

			if response.StatusCode >= 300 {
				return fmt.Errorf("unexpected status code: %d", response.StatusCode)
			}

			return PrintWallet(cmd.OutOrStdout(), response.GetWalletResponse.Data)
		}),
	)
}

func PrintWallet(out io.Writer, wallet shared.WalletWithBalances) error {
	fctl.Section.Println("Information")
	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("ID"), fmt.Sprint(wallet.ID)})
	tableData = append(tableData, []string{pterm.LightCyan("Name"), wallet.Name})

	if err := pterm.DefaultTable.
		WithWriter(out).
		WithData(tableData).
		Render(); err != nil {
		return err
	}

	fctl.Section.Println("Balances")
	if len(wallet.Balances.Main.Assets) == 0 {
		fctl.Println("No balances found.")
	} else {
		tableData = pterm.TableData{}
		tableData = append(tableData, []string{"Asset", "Amount"})
		for asset, amount := range wallet.Balances.Main.Assets {
			tableData = append(tableData, []string{asset, fmt.Sprint(amount)})
		}
		if err := pterm.DefaultTable.
			WithHasHeader(true).
			WithWriter(out).
			WithData(tableData).
			Render(); err != nil {
			return err
		}
	}

	return fctl.PrintMetadata(out, wallet.Metadata)
}
