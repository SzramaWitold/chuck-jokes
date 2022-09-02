package api

func (s *Server) setRoutes() {
	s.Engine.POST("/login", s.Controller.User.Login())
	s.Engine.POST("/register", s.Controller.User.Register())

	s.Engine.GET("/jokeoftheday", s.Controller.Joke.GetJokeOfADay())
	s.Engine.GET("/jokes", s.Controller.Joke.GetAll())
	s.Engine.GET("/jokes/:ID", s.Controller.Joke.Get())
	s.Engine.GET("/categories/:ID", s.Controller.Category.GetCategory())

	restricted := s.Engine.Group("/").Use(s.Middleware.Auth.Auth)
	restricted.POST("/me", s.Controller.User.GetMe())
	restricted.PUT("/favourite", s.Controller.User.AddFavourite())
	restricted.GET("/favourite", s.Controller.User.GetFavourites())
	restricted.POST("/categories", s.Controller.Category.CreateCategory())
	restricted.GET("/jokes/:ID/statistic", s.Controller.Joke.GetStatistic())
	restricted.PUT("/categories/:ID", s.Controller.Category.AddToCategory())
	restricted.DELETE("/categories/:ID", s.Controller.Category.RemoveFromCategory())
	restricted.PUT("/categories/:ID/access", s.Controller.Category.SetAccessCategory())
}
