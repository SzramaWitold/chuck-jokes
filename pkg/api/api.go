package api

import (
	_ "chuck-jokes/docs"
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server base gin server
type Server struct {
	Engine     *gin.Engine
	Controller controllers.ControllerHandler
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
	server.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return server
}
