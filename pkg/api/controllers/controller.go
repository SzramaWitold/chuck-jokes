package controllers

import (
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/token"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory() func(c *gin.Context)
	AddToCategory() func(c *gin.Context)
	SetAccessCategory() func(c *gin.Context)
	RemoveFromCategory() func(c *gin.Context)
	GetCategory() func(c *gin.Context)
}

type UserController interface {
	GetMe() func(c *gin.Context)
	Login() func(c *gin.Context)
	GetFavourites() func(c *gin.Context)
	AddFavourite() func(c *gin.Context)
	Register() func(c *gin.Context)
}

type JokeController interface {
	GetJokes() func(c *gin.Context)
	GetJokeOfADay() func(c *gin.Context)
	GetJoke() func(c *gin.Context)
	GetStatistic() func(c *gin.Context)
}

type ControllerHandler interface {
	CategoryController
	UserController
	JokeController
}

type Controller struct {
	Request    requests.RequestHandler
	Response   responses.ResponseHandler
	JWT        *token.TokenHandler
	Repository *repositories.Repository
}

func NewController(
	jwt *token.TokenHandler,
	request requests.RequestHandler,
	response responses.ResponseHandler,
	repository *repositories.Repository) *Controller {
	return &Controller{
		Request:    request,
		Response:   response,
		JWT:        jwt,
		Repository: repository,
	}
}
