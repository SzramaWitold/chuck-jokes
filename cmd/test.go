package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "cmd tests",
	Long:  `Test work in progress functionality`,
	Run: func(_ *cobra.Command, _ []string) {
		seeder := database.NewSeeder(di.GORM())
		seeder.Seed()
	},
}
