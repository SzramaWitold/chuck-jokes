package controllers

import (
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/token"
	"github.com/gin-gonic/gin"
)

type ICategory interface {
	CreateCategory() func(c *gin.Context)
	AddToCategory() func(c *gin.Context)
	SetAccessCategory() func(c *gin.Context)
	RemoveFromCategory() func(c *gin.Context)
	GetCategory() func(c *gin.Context)
}

type IUser interface {
	GetMe() func(c *gin.Context)
	Login() func(c *gin.Context)
	GetFavourites() func(c *gin.Context)
	AddFavourite() func(c *gin.Context)
	Register() func(c *gin.Context)
}

type IJoke interface {
	GetJokes() func(c *gin.Context)
	GetJokeOfADay() func(c *gin.Context)
	GetJoke() func(c *gin.Context)
	GetStatistic() func(c *gin.Context)
}

type IController interface {
	ICategory
	IUser
	IJoke
}

type Controller struct {
	Request    requests.IRequest
	Response   responses.IResponse
	JWT        *token.IHandler
	Repository *repositories.Repository
}

func NewController(
	jwt *token.IHandler,
	request requests.IRequest,
	response responses.IResponse,
	repository *repositories.Repository) *Controller {
	return &Controller{
		Request:    request,
		Response:   response,
		JWT:        jwt,
		Repository: repository,
	}
}
