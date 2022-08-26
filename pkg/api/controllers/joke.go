package controllers

import (
	"chuck-jokes/pkg/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetJokes godoc
// @Summary      GetJokes
// @Description  Get list of jokes
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.PaginateJokes
// @Failure      401  {object}  responses.Error
// @Router       /jokes [get]
func (cont *Controller) GetJokes() func(c *gin.Context) {
	return func(c *gin.Context) {
		pagRequest := cont.Request.NewPagination(c)
		repJokes := cont.Repository.Joke.GetJokes(pagRequest.Page, pagRequest.PerPage)

		c.JSON(http.StatusOK, cont.Response.NewPaginateJokes(repJokes))
	}
}

// GetJoke godoc
// @Summary      GetJoke
// @Description  get specify joke
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.Joke
// @Failure      401  {array}  responses.Error
// @Router       /jokes/{ID} [get]
func (cont *Controller) GetJoke() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewJoke(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		repJoke := cont.Repository.Joke.GetJoke(request.JokeID)

		c.JSON(http.StatusOK, cont.Response.NewJoke(repJoke))
	}
}

// GetStatistic godoc
// @Summary      GetStatistic
// @Description  get statistic of the specify joke
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.JokeStatistic
// @Failure      401  {array}  responses.Error
// @Router       /jokes/{ID}/statistic [get]
func (cont *Controller) GetStatistic() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewJoke(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		repJoke, favNumber := cont.Repository.Joke.GetStatistic(request.JokeID)

		c.JSON(http.StatusOK, cont.Response.NewJokeStatistic(repJoke, favNumber))
	}
}

// GetJokeOfADay godoc
// @Summary      GetJokeOfADay
// @Description  get current joke of a day
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.Joke
// @Failure      401  {array}  responses.Error
// @Router       /jokeoftheday [get]
func (cont *Controller) GetJokeOfADay() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewJokeOfADay(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		if request.Date == "" {
			request.Date = utilities.GetToday().String()
		}

		joke := cont.Repository.Joke.JokeOfTheDay(request.Date)

		if joke == nil {
			c.JSON(http.StatusNotFound, cont.Response.NewError(fmt.Errorf("can not find joke for date: %v", request.Date)))
			return
		}
		c.JSON(http.StatusOK, cont.Response.NewJoke(joke))
	}
}
