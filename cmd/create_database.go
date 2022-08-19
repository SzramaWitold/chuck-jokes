package cmd

import (
	"chuck-jokes/pkg/database/gorm"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createDatabaseCmd)
}

var createDatabaseCmd = &cobra.Command{
	Use:   "database:create",
	Short: "Create new database",
	Long:  `Create new database based on .env file`,
	Run: func(_ *cobra.Command, _ []string) {
		gorm.CreateDatabase(
			os.Getenv("DB_TYPE"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"))
	},
}
