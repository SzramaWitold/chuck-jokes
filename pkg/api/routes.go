package api

func (s *Server) setRoutes() {
	s.Engine.POST("/login", s.Controller.Login())
	s.Engine.POST("/register", s.Controller.Register())

	s.Engine.GET("/jokeoftheday", s.Controller.GetJokeOfADay())
	s.Engine.GET("/jokes", s.Controller.GetJokes())

	restricted := s.Engine.Group("/").Use(s.Middleware.Auth.Auth)
	restricted.POST("/me", s.Controller.GetMe())
	restricted.PUT("/favourite", s.Controller.AddFavourite())
	restricted.GET("/favourite", s.Controller.GetFavourites())
	restricted.POST("/categories", s.Controller.CreateCategory())
	restricted.PUT("/categories/:ID", s.Controller.AddToCategory())
}
