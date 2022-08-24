package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database/gorm"
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
		seeder := gorm.NewSeeder(di.GORM())
		seeder.Seed()
	},
}
