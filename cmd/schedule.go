package cmd

import (
	"chuck-jokes/pkg/crone"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scheduler)
}

var scheduler = &cobra.Command{
	Use:   "schedule:run",
	Short: "schedul all job inside crone",
	Long:  `Schedule everything inside crone package`,
	Run: func(_ *cobra.Command, _ []string) {
		crone.Schedule(false)
	},
}
