package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tanan/wg-config-generator/config"
)

var rootCmd = &cobra.Command{
	Use:   "wgconf",
	Short: "Wireguard Config Generator",
	Long:  `wgconf is a CLI to generate server and client configrations.`,
}

func init() {
	rootCmd.PersistentFlags().StringP("configFile", "c", "config.yaml", "config file path")
}

func initConfig(cmd *cobra.Command) {
	configFile, _ := cmd.Flags().GetString("configFile")
	config.LoadConfig(configFile)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
