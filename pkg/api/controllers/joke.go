package controllers

import (
	"chuck-jokes/pkg/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cont *Controller) GetFavourites() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewFavourites(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(requestErr))
		}
		jokes := cont.Repository.Joke.GetFavourites(request.Page, request.PerPage, request.UserID)

		c.JSON(http.StatusOK, cont.Response.PaginateJokes(jokes))
	}
}

func (cont *Controller) AddFavourite() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewAddFavouriteRequest(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		} else {
			repErr := cont.Repository.User.AddFavourite(request.UserID, request.JokeID)

			if repErr != nil {
				c.JSON(http.StatusExpectationFailed, repErr.Error())
				return
			}
			c.JSON(http.StatusOK, cont.Response.NewSuccess("success"))
		}
	}
}

func (cont *Controller) GetJokes() func(c *gin.Context) {
	return func(c *gin.Context) {
		pagRequest := cont.Request.NewPagination(c)
		repJokes := cont.Repository.Joke.GetJokes(pagRequest.Page, pagRequest.PerPage)

		c.JSON(http.StatusOK, cont.Response.PaginateJokes(repJokes))
	}
}

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
