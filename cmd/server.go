package cmd

import (
	"fmt"
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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initConfig(cmd)

		action := args[0]
		if action != "create" {
			slog.Error(fmt.Sprintf("unknown command %v", action))
		}

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
