package api

import (
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoutes(engine gin.Engine) {
	engine.GET("/jokeoftheday", jokeoftheday)
	engine.GET("/jokes", getJokes)
}

func getJokes(c *gin.Context) {
	jokes := repositories.GetJokes()

	c.JSON(http.StatusOK, jokes)
}
func jokeoftheday(c *gin.Context) {
	joke := repositories.JokeOfTheDay(utilities.GetToday().String())
	if joke == nil {
		joke = repositories.JokeOfTheDay(utilities.GetYesterday().String())
	}
	c.JSON(http.StatusOK, joke)
}
