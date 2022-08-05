package api

import (
	"chuck-jokes/pkg/api/controllers"
	"chuck-jokes/pkg/api/middlewares"
)

func (s *Server) setRoutes() {
	controller := controllers.NewController(s.DB)
	s.Engine.POST("/login", controller.Login())
	s.Engine.Use(middlewares.Auth).POST("/me", controller.GetMe())
	s.Engine.Use(middlewares.Auth).PUT("/favourite", controller.AddFavourite())
	s.Engine.Use(middlewares.Auth).GET("/favourite", controller.GetFavourites())
	s.Engine.Use(middlewares.Auth).POST("/categories", controller.CreateCategory())
	s.Engine.Use(middlewares.Auth).PUT("/categories/:id", controller.CreateCategory())

	s.Engine.GET("/jokeoftheday", controller.GetJokeOfADay())
	s.Engine.GET("/jokes", controller.GetJokes())
}
