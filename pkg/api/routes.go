package api

import (
	"chuck-jokes/pkg/api/middlewares"
	"chuck-jokes/pkg/api/requests"
	"chuck-jokes/pkg/api/responses"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/token"
	"chuck-jokes/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) setRoutes() {
	s.Engine.POST("/login", s.login)
	s.Engine.Use(middlewares.Auth).POST("/me", s.me)
	s.Engine.Use(middlewares.Auth).PUT("/favourite", s.addFavourite)
	s.Engine.Use(middlewares.Auth).GET("/favourite", s.getFavourite)

	s.Engine.GET("/jokeoftheday", s.jokeOfTheDay)
	s.Engine.GET("/jokes", s.getJokes)
}

func (s *Server) getFavourite(c *gin.Context) {
	repository := repositories.NewJokeRepository(s.db)
	request, requestErr := requests.NewFavouritesRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusBadRequest, requestErr.Error())
	}
	jokes := repository.GetFavourites(request.Page, request.PerPage, request.UserID)

	c.JSON(http.StatusOK, jokes)
}

func (s *Server) addFavourite(c *gin.Context) {
	request, requestErr := requests.NewAddFavouriteRequest(c)

	if requestErr != nil {
		c.JSON(http.StatusBadRequest, requestErr.Error())
	} else {
		repository := repositories.NewUserRepository(s.db)
		repErr := repository.AddFavourite(request.UserID, request.JokeID)

		if repErr != nil {
			c.JSON(http.StatusExpectationFailed, repErr.Error())
			return
		}
		c.JSON(http.StatusOK, "success")
	}
}

func (s *Server) me(c *gin.Context) {
	c.JSON(http.StatusOK, "I am in")
}

func (s *Server) login(c *gin.Context) {
	userRepository := repositories.NewUserRepository(s.db)
	user := userRepository.Authenticate("test", "test")
	if user != nil {
		validator := token.NewValidator()

		c.JSON(http.StatusOK, responses.NewTokenResponse(validator.CreateToken(user)))
	} else {
		c.JSON(http.StatusUnauthorized, "Wrong credentials!")
	}
}

func (s *Server) getJokes(c *gin.Context) {
	jokeRepository := repositories.NewJokeRepository(s.db)
	pagRequest := requests.NewPaginationRequest(c)
	jokes := jokeRepository.GetJokes(pagRequest.Page, pagRequest.PerPage)

	c.JSON(http.StatusOK, jokes)
}

func (s *Server) jokeOfTheDay(c *gin.Context) {
	jokeRepository := repositories.NewJokeRepository(s.db)
	joke := jokeRepository.JokeOfTheDay(utilities.GetToday().String())
	if joke == nil {
		joke = jokeRepository.JokeOfTheDay(utilities.GetYesterday().String())
	}
	c.JSON(http.StatusOK, joke)
}
