package api

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
}

func StartEngine() Server {
	server := Server{
		Engine: gin.Default(),
	}
	setRoutes(*server.Engine)
	
	return server
}
