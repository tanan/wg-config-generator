package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/handler"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a server configuration",
	Long: `Create a server configuration with given parameters. For example:

wgconf server create -c /path/to/config.yaml -o /path/to/output.conf
`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig(cmd)

		h := handler.NewHandler(handler.NewCommand(), config.GetConfig())
		peers, err := h.GetClientList()
		if err != nil {
			slog.Error("failed to get client list", slog.String("error", err.Error()))
		}
		serverConfig, err := h.CreateServerConfig()
		if err != nil {
			slog.Error("failed to create server config", slog.String("error", err.Error()))
		}
		if err := h.WriteServerConfig(serverConfig, peers); err != nil {
			slog.Error("failed to write server config", slog.String("error", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
