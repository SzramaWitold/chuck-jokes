package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/api"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create and get server",
	Long:  `Create and get server for chuck noris jokes aplication`,
	Run: func(_ *cobra.Command, _ []string) {
		server := api.StartEngine(di.GORM())
		serverError := server.Engine.Run(":8080")
		if serverError != nil {
			log.Panic(serverError)
		}
	},
}

// Execute default execution for cmd
func Execute() {
	if cmdError := rootCmd.Execute(); cmdError != nil {
		log.Panic(cmdError)
	}
}
