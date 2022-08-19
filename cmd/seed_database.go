package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database/gorm"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "database:seed",
	Short: "seed database with fake data",
	Long:  `Seed data based on fake models from database`,
	Run: func(_ *cobra.Command, _ []string) {
		seeder := gorm.NewSeeder(di.GORM(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		))
		seeder.Seed()
	},
}
