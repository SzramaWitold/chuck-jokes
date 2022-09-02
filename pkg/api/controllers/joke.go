package controllers

import (
	"fmt"
	"net/http"

	"chuck-jokes/pkg/repositories/gorm"

	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/utilities"

	"github.com/gin-gonic/gin"
)

type JokeController interface {
	GetAll() func(c *gin.Context)
	GetJokeOfADay() func(c *gin.Context)
	Get() func(c *gin.Context)
	GetStatistic() func(c *gin.Context)
}

type Joke struct {
	request    requests.RequestHandler
	response   responses.ResponseHandler
	repository *gorm.Repository
}

func NewJoke(request requests.RequestHandler, response responses.ResponseHandler, repository *gorm.Repository) *Joke {
	return &Joke{request: request, response: response, repository: repository}
}

// GetAll godoc
// @Summary      GetJokes
// @Description  Get list of jokes
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.PaginateJokes
// @Failure      401  {object}  responses.Error
// @Router       /jokes [get]
func (j *Joke) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		pageRequest := j.request.NewFindCollection(c)
		repJokes := j.repository.Joke.FindAll(pageRequest)

		c.JSON(http.StatusOK, j.response.NewPaginateJokes(repJokes))
	}
}

// Get godoc
// @Summary      GetJoke
// @Description  get specify joke
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.Joke
// @Failure      401  {object}  responses.Error
// @Router       /jokes/{ID} [get]
func (j *Joke) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := j.request.NewJoke(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, j.response.NewError(requestErr))

			return
		}

		repJoke := j.repository.Joke.Find(request.JokeID)

		if addStatisticErr := j.repository.JokeStatistic.AddShowByJokeID(request.JokeID); addStatisticErr != nil {
			c.JSON(http.StatusBadRequest, j.response.NewError(addStatisticErr))

			return
		}

		c.JSON(http.StatusOK, j.response.NewJoke(repJoke))
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
// @Failure      401  {object}  responses.Error
// @Router       /jokes/{ID}/statistic [get]
func (j *Joke) GetStatistic() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := j.request.NewJoke(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, j.response.NewError(requestErr))

			return
		}

		repJoke, favNumber, jokeStatisticErr := j.repository.JokeStatistic.FindByJokeID(request.JokeID)

		if jokeStatisticErr != nil {
			c.JSON(http.StatusBadRequest, j.response.NewError(jokeStatisticErr))

			return
		}

		c.JSON(http.StatusOK, j.response.NewJokeStatistic(repJoke, favNumber))
	}
}

// GetJokeOfADay godoc
// @Summary      GetJokeOfADay
// @Description  get current joke of a day
// @Tags         Joke
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.Joke
// @Failure      401  {object}  responses.Error
// @Router       /jokeoftheday [get]
func (j *Joke) GetJokeOfADay() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := j.request.NewJokeOfADay(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, j.response.NewError(requestErr))

			return
		}

		if request.Date == "" {
			request.Date = utilities.GetToday().String()
		}

		joke := j.repository.Joke.JokeOfTheDay(request.Date)

		if joke == nil {
			c.JSON(http.StatusNotFound, j.response.NewError(fmt.Errorf("can not find joke for date: %v", request.Date)))

			return
		}

		c.JSON(http.StatusOK, j.response.NewJoke(joke))
	}
}
