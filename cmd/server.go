package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a server configuration",
	Long: `Create a server configuration with given parameters. For example:

wgconf server create -c /path/to/config.yaml -o /path/to/output.conf
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
