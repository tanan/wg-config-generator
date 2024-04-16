/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
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

		user := args[0]
		fmt.Println("user:", user)
		fmt.Println(config.GetConfig().WorkDir)
	},
}

func init() {
	clientCmd.AddCommand(createClientCmd)
	rootCmd.AddCommand(clientCmd)
}
