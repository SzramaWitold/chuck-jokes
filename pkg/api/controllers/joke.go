package controllers

import (
	"chuck-jokes/pkg/api/requests"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cont *Controller) GetFavourites() func(c *gin.Context) {
	return func(c *gin.Context) {
		repository := repositories.NewJoke(cont.DB)
		request, requestErr := requests.NewFavourites(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, requestErr.Error())
		}
		jokes := repository.GetFavourites(request.Page, request.PerPage, request.UserID)

		c.JSON(http.StatusOK, jokes)
	}
}

func (cont *Controller) AddFavourite() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := requests.NewAddFavouriteRequest(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, requestErr.Error())
		} else {
			repository := repositories.NewUser(cont.DB)
			repErr := repository.AddFavourite(request.UserID, request.JokeID)

			if repErr != nil {
				c.JSON(http.StatusExpectationFailed, repErr.Error())
				return
			}
			c.JSON(http.StatusOK, "success")
		}
	}
}

func (cont *Controller) GetJokes() func(c *gin.Context) {
	return func(c *gin.Context) {
		jokeRepository := repositories.NewJoke(cont.DB)
		pagRequest := requests.NewPagination(c)
		jokes := jokeRepository.GetJokes(pagRequest.Page, pagRequest.PerPage)

		c.JSON(http.StatusOK, jokes)
	}
}

func (cont *Controller) GetJokeOfADay() func(c *gin.Context) {
	return func(c *gin.Context) {
		jokeRepository := repositories.NewJoke(cont.DB)
		joke := jokeRepository.JokeOfTheDay(utilities.GetToday().String())
		if joke == nil {
			joke = jokeRepository.JokeOfTheDay(utilities.GetYesterday().String())
		}
		c.JSON(http.StatusOK, joke)
	}
}
