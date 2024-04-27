/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/handler"
	"github.com/tanan/wg-config-generator/model"
	"github.com/tanan/wg-config-generator/util"
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
		t, _ := cmd.Flags().GetString("output-type")
		if err := model.OutputType(t).Valid(); err != nil {
			slog.Error("flag `output-type` error", slog.String("error", err.Error()))
			os.Exit(1)
		}

		h := handler.NewHandler(handler.NewCommand(), config.GetConfig())

		name := args[0]
		address, _ := cmd.Flags().GetString("ip-address")

		clientConfig, err := h.CreateClientConfig(name, address)
		if err != nil {
			var fileError util.FileError
			var keyError util.KeyError
			if errors.As(err, &keyError) {
				slog.Error("KeyError when func handler.CreateClientConfig()", slog.String("error", err.Error()))
			} else if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.CreateClientConfig()", slog.String("path", fileError.Path))
			} else {
				slog.Error("Error when func handler.CreateClientConfig()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}

		if err := h.SaveClientSetting(clientConfig); err != nil {
			var fileError util.FileError
			if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.SaveClientSetting()", slog.String("path", fileError.Path))
			} else {
				slog.Error("Error when func handler.SaveClientSetting()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}

		serverConfig, err := h.CreateServerConfig()
		if err != nil {
			var fileError util.FileError
			if errors.As(err, &fileError) {
				slog.Error("FileError when func handler.CreateServerConfig()", slog.String("path", fileError.Path))
			} else {
				slog.Error("Error when func handler.CreateServerConfig()", slog.String("error", err.Error()))
			}
			os.Exit(1)
		}

		var outputErr error
		switch model.OutputType(t) {
		case model.Text:
			outputErr = h.WriteClientConfig(clientConfig, serverConfig)
		case model.Email:
			outputErr = h.SendClientConfigByEmail(clientConfig, serverConfig)
		default:
			slog.Info("this output type is not implemented yet", slog.String("type", t))
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
