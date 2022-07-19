package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		database.BaseSeed()
	},
}
