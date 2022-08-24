package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/api"
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/api/middlewares"
	"chuck-jokes/pkg/repositories"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create and get server",
	Long:  `Create and get server for chuck noris jokes application`,
	Run: func(_ *cobra.Command, _ []string) {
		gorm := di.GORM()
		jwt := di.JWT()
		request := requests.NewRequestValidator(di.VALIDATOR(gorm))
		response := responses.NewResponse()
		repository := repositories.NewRepository(gorm)
		controller := controllers.NewController(jwt, request, response, repository)
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
