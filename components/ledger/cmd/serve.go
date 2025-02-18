package cmd

import (
	"github.com/formancehq/ledger/pkg/api/middlewares"
	"github.com/formancehq/stack/libs/go-libs/ballast"
	"github.com/formancehq/stack/libs/go-libs/httpserver"
	app "github.com/formancehq/stack/libs/go-libs/service"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	queryLimitReadLogsFlag = "query-limit-read-logs"
	ballastSizeInBytesFlag = "ballast-size"
	numscriptCacheMaxCount = "numscript-cache-max-count"
)

func NewServe() *cobra.Command {
	cmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.New(cmd.OutOrStdout(), resolveOptions(
				cmd.OutOrStdout(),
				ballast.Module(viper.GetUint(ballastSizeInBytesFlag)),
				fx.Invoke(func(lc fx.Lifecycle, h chi.Router) {

					if viper.GetBool(app.DebugFlag) {
						wrappedRouter := chi.NewRouter()
						wrappedRouter.Use(middlewares.Log())
						wrappedRouter.Mount("/", h)
						h = wrappedRouter
					}

					lc.Append(httpserver.NewHook(viper.GetString(bindFlag), h))
				}),
			)...).Run(cmd.Context())
		},
	}
	cmd.Flags().Int(queryLimitReadLogsFlag, 10000, "Query limit read logs")
	cmd.Flags().Uint(ballastSizeInBytesFlag, 0, "Ballast size in bytes, default to 0")
	cmd.Flags().Int(numscriptCacheMaxCount, 1024, "Numscript cache max count")
	return cmd
}
