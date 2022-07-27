package api

import (
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/utilities"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server)  setRoutes() {
	s.Engine.GET("/jokeoftheday", s.jokeoftheday)
	s.Engine.GET("/jokes", s.getJokes)
}

func (s *Server) getJokes(c *gin.Context) {
	jokeRepository := repositories.NewJokeRepository(s.db)
	page, perPage := getPaginationSetup(c)
	jokes := jokeRepository.GetJokes(page, perPage)

	c.JSON(http.StatusOK, jokes)
}

func (s *Server) jokeoftheday(c *gin.Context) {
	jokeRepository := repositories.NewJokeRepository(s.db)
	joke := jokeRepository.JokeOfTheDay(utilities.GetToday().String())
	if joke == nil {
		joke = jokeRepository.JokeOfTheDay(utilities.GetYesterday().String())
	}
	c.JSON(http.StatusOK, joke)
}

func getPaginationSetup(c *gin.Context) (int, int) {
	query := c.Request.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		log.Println("Wrong type provide as a page parameter")
	}

	perPage, err := strconv.Atoi(query.Get("per_page"))

	if err != nil {
		log.Println("Wrong type provide as a per_page parameter")
	}
	return page, perPage
}
