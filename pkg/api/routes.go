package api

func (s *Server) setRoutes() {
	s.Engine.POST("/login", s.Controller.Login())
	s.Engine.Use(s.Middleware.Auth).POST("/me", s.Controller.GetMe())
	s.Engine.Use(s.Middleware.Auth).PUT("/favourite", s.Controller.AddFavourite())
	s.Engine.Use(s.Middleware.Auth).GET("/favourite", s.Controller.GetFavourites())
	s.Engine.Use(s.Middleware.Auth).POST("/categories", s.Controller.CreateCategory())
	s.Engine.Use(s.Middleware.Auth).PUT("/categories/:ID", s.Controller.AddToCategory())

	s.Engine.GET("/jokeoftheday", s.Controller.GetJokeOfADay())
	s.Engine.GET("/jokes", s.Controller.GetJokes())
}
