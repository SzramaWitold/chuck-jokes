package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/api"
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/api/middlewares"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create and get server",
	Long:  `Create and get server for chuck noris jokes application`,
	Run: func(_ *cobra.Command, _ []string) {
		gorm := di.GORM(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		jwt := di.JWT(
			os.Getenv("SECRET"),
			os.Getenv("TTL"), os.Getenv("REFRESH_TTL"))
		request := requests.NewRequest(di.VALIDATOR())
		response := responses.NewResponse()
		controller := controllers.NewController(gorm, jwt, request, response)
		middleware := middlewares.NewMiddleware(jwt)

		server := api.StartEngine(controller, middleware)

		serverError := server.Engine.Run(":8080")
		if serverError != nil {
			log.Panic(serverError)
		}
	},
}

// Execute default execution for cmd
func Execute() {
	if cmdError := rootCmd.Execute(); cmdError != nil {
		log.Panic(cmdError)
	}
}
