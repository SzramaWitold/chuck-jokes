package cmd

import (
	"chuck-jokes/di"
	gorm2 "chuck-jokes/pkg/database/gorm"
	"chuck-jokes/pkg/database/gorm/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateDatabaseCmd)
}

var migrateDatabaseCmd = &cobra.Command{
	Use:   "database:migrate",
	Short: "migrate database",
	Long:  `Migrate database based on database.GetAllModels function`,
	Run: func(_ *cobra.Command, _ []string) {
		dbManager := gorm2.NewManager(di.GORM())
		dbManager.MigrateModels(models.GetAllModels()...)
	},
}
