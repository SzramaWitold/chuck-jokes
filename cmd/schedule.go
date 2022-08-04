package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/crone"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(schedulerCmd)
}

var schedulerCmd = &cobra.Command{
	Use:   "schedule:run",
	Short: "schedule all job inside crone",
	Long:  `Schedule everything inside crone package`,
	Run: func(_ *cobra.Command, _ []string) {
		scheduler := crone.NewCronScheduler(di.Scheduler())
		scheduler.Schedule(false)
	},
}
