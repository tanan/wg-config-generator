/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/handler"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Wireguard client config command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
	},
}

var createClientCmd = &cobra.Command{
	Use:   "create",
	Short: "Create client config with a given user name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// initialize config
		initConfig(cmd)
		h := handler.NewHandler(handler.NewCommand(), config.GetConfig())

		name := args[0]
		address, _ := cmd.Flags().GetString("ip-address")

		clientConfig, err := h.CreateClientConfig(name, address)
		if err != nil {
			slog.Error("output error", slog.String("error", err.Error()))
			os.Exit(1)
		}

		outputType, _ := cmd.Flags().GetString("output-type")
		var outputErr error
		switch outputType {
		case "text":
			outputErr = h.WriteClientConfig(clientConfig)
		case "email":
			outputErr = h.SendClientConfigByEmail(clientConfig)
		default:
			outputErr = fmt.Errorf("output type %s is not found", outputType)
		}
		if outputErr != nil {
			slog.Error("output error", slog.String("error", outputErr.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	createClientCmd.Flags().StringP("ip-address", "i", "", "ip address for client")
	createClientCmd.Flags().StringP("output-type", "t", "text", "client configuration output type: [text|email]")
	createClientCmd.MarkFlagRequired("ip-address")
	createClientCmd.MarkFlagRequired("output-type")
	clientCmd.AddCommand(createClientCmd)
	rootCmd.AddCommand(clientCmd)
}
