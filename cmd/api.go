package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create and get server",
	Long:  `Create and get server for chuck noris jokes aplication`,
	Run: func(_ *cobra.Command, _ []string) {
		server := api.StartEngine(di.GORM())
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
