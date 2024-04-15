package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a server configuration",
	Long: `Create a server configuration with given params. For example:

wgconf server create --config /path/to/config.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// TODO: 必要な引数
	// config: path to config file

	// contents of the config file
	// client file foloder
	// endpoint / port / interface / ip_address / pubkey / privkey(another file path) / presharedkey(another file path)

}
