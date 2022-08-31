package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database/gorm"
	gormRepository "chuck-jokes/pkg/repositories/gorm"
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
		repository := gormRepository.NewRepository(di.GORM())
		seeder := gorm.NewSeeder(repository)
		seeder.Seed()
	},
}
