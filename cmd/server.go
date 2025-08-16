package cmd

import (
	"math"

	"github.com/evertontomalok/go-rest-sample/internal/adapters/infra"
	"github.com/evertontomalok/go-rest-sample/internal/app"
	"github.com/evertontomalok/go-rest-sample/internal/app/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		config := app.Configure(ctx)
		repository := infra.NewMemDB(
			infra.WithMaxSize(int(math.Pow10(10))), // 10 ** 5 == 100.000
		)
		server.RunServer(ctx, config, repository)
	},
}

func init() {
	serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(serverCmd)
}
