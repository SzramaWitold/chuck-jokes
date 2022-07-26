package api

import (
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/utilities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setRoutes(engine gin.Engine) {
	engine.GET("/jokeoftheday", jokeoftheday)
	engine.GET("/jokes", getJokes)
}

func getJokes(c *gin.Context) {
	page, perPage := getPaginationSetup(c)
	jokes := repositories.GetJokes(page, perPage)

	c.JSON(http.StatusOK, jokes)
}

func jokeoftheday(c *gin.Context) {
	joke := repositories.JokeOfTheDay(utilities.GetToday().String())
	if joke == nil {
		joke = repositories.JokeOfTheDay(utilities.GetYesterday().String())
	}
	c.JSON(http.StatusOK, joke)
}

func getPaginationSetup(c *gin.Context) (int, int) {
	query := c.Request.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		fmt.Println("Wrong type provide as a page parameter")
	}

	perPage, err := strconv.Atoi(query.Get("per_page"))

	if err != nil {
		fmt.Println("Wrong type provide as a per_page parameter")
	}
	return page, perPage
}
