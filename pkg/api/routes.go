package api

import (
	"chuck-jokes/pkg/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoutes(engine gin.Engine) {
	engine.GET("/jokeoftheday", jokeoftheday)
}

func jokeoftheday(c *gin.Context) {
	joke := requests.CallRandom()
	c.JSON(http.StatusOK, joke)
}
