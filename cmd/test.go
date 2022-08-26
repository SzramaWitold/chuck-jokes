package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "...",
	Long:  `....`,
	Run: func(_ *cobra.Command, _ []string) {
		num := -4

		println(uint(num))
	},
}
