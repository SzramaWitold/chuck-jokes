package api

import "github.com/gin-gonic/gin"

// Server base gin server
type Server struct {
	Engine *gin.Engine
}

// StartEngine start gin engine, add routes and return server struct
func StartEngine() Server {
	server := Server{
		Engine: gin.Default(),
	}
	setRoutes(*server.Engine)

	return server
}
