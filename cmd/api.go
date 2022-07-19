package cmd

import (
	"chuck-jokes/pkg/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create and get server",
	Long:  `Create and get server for chuck noris jokes aplication`,
	Run: func(cmd *cobra.Command, args []string) {
		server := api.StartEngine()
		server.Engine.Run(":8080")
	},
}

// Execute default execution for cmd 
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
