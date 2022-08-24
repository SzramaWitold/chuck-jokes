package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/validator"
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
		gorm := di.GORM()
		val := validator.NewValidator(gorm)
		fmt.Println(val.Validate(requests.AddFavourite{}, map[string]string{
			"UserID": "4",
			"JokeID": "4",
			"Test":   "sd",
		}))
	},
}
