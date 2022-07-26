package cmd

import (
	"chuck-jokes/pkg/database"
	"chuck-jokes/pkg/repositories"
	"fmt"

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
		var joke = database.Joke{}
		joke.JokeResponse.ID = "8q7JD4FFRhqS-iiWqZkHsg"
		fmt.Println(repositories.JokeExistInLastMonth(&joke))
	},
}
