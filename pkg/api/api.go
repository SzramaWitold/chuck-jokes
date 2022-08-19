package api

import (
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/middlewares"
	"github.com/gin-gonic/gin"
)

// Server base gin server
type Server struct {
	Engine     *gin.Engine
	Controller controllers.IController
	Middleware *middlewares.Middleware
}

// StartEngine start gin engine, add routes and return server struct
func StartEngine(controller *controllers.Controller, middleware *middlewares.Middleware) Server {
	server := Server{
		Engine:     gin.Default(),
		Controller: controller,
		Middleware: middleware,
	}
	server.setRoutes()

	return server
}
