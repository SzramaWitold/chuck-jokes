package api

import (
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/middlewares"
	"chuck-jokes/pkg/token"
	"chuck-jokes/pkg/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server base gin server
type Server struct {
	Engine     *gin.Engine
	Controller *controllers.Controller
	Middleware *middlewares.Middleware
}

// StartEngine start gin engine, add routes and return server struct
func StartEngine(db *gorm.DB, validator *validator.Validator, jwt *token.Handler) Server {
	server := Server{
		Engine:     gin.Default(),
		Controller: controllers.NewController(db, validator),
		Middleware: middlewares.NewMiddleware(jwt),
	}
	server.setRoutes()

	return server
}
