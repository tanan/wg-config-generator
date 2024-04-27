package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/handler"
	"github.com/tanan/wg-config-generator/util"
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
			os.Exit(1)
		}

		h := handler.NewHandler(handler.NewCommand(), config.GetConfig())
		peers, err := h.GetClientList()
		if err != nil {
			var fileError util.FileError
			if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.GetClientList()", slog.String("path", fileError.Path), slog.String("error", err.Error()))
			} else {
				slog.Error("Error when func handler.GetClientList()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}
		serverConfig, err := h.CreateServerConfig()
		if err != nil {
			var fileError util.FileError
			if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.CreateServerConfig()", slog.String("path", fileError.Path), slog.String("error", err.Error()))
			} else {
				slog.Error("Error when func handler.CreateServerConfig()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}
		if err := h.WriteServerConfig(serverConfig, peers); err != nil {
			var fileError util.FileError
			if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.WriteServerConfig()", slog.String("path", fileError.Path), slog.String("error", err.Error()))
			} else {
				slog.Error("Error when func handler.WriteServerConfig()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
