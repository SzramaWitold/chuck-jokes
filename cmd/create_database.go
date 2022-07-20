package cmd

import (
	"chuck-jokes/pkg/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createDatabaseCmd)
}

var createDatabaseCmd = &cobra.Command{
	Use:   "database:create",
	Short: "Create new database",
	Long:  `Create new database based on .env file name`,
	Run: func(_ *cobra.Command, _ []string) {
		database.CreateDatabase()
	},
}
