package cmd

import (
	"chuck-jokes/di"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(schedulerCmd)
}

var schedulerCmd = &cobra.Command{
	Use:   "schedule:run",
	Short: "schedul all job inside crone",
	Long:  `Schedule everything inside crone package`,
	Run: func(_ *cobra.Command, _ []string) {
		scheduler := di.CronScheduler()
		scheduler.Schedule(false)
	},
}
