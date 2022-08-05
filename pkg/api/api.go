package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server base gin server
type Server struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

// StartEngine start gin engine, add routes and return server struct
func StartEngine(db *gorm.DB) Server {
	server := Server{
		Engine: gin.Default(),
		DB:     db,
	}
	server.setRoutes()

	return server
}
