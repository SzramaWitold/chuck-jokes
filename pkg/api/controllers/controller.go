package controllers

import (
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/token"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type IController interface {
	CreateCategory() func(c *gin.Context)
	AddToCategory() func(c *gin.Context)
	GetMe() func(c *gin.Context)
	Login() func(c *gin.Context)
	GetFavourites() func(c *gin.Context)
	AddFavourite() func(c *gin.Context)
	GetJokes() func(c *gin.Context)
	GetJokeOfADay() func(c *gin.Context)
	Register() func(c *gin.Context)
	SetAccessCategory() func(c *gin.Context)
	GetCategory() func(c *gin.Context)
	RemoveFromCategory() func(c *gin.Context)
	GetJoke() func(c *gin.Context)
	GetStatistic() func(c *gin.Context)
}

type Controller struct {
	DB       *gorm.DB
	Request  requests.IRequest
	Response responses.IResponse
	JWT      *token.IHandler
}

func NewController(
	db *gorm.DB,
	jwt *token.IHandler,
	request requests.IRequest,
	response responses.IResponse) *Controller {
	return &Controller{
		DB:       db,
		Request:  request,
		Response: response,
		JWT:      jwt,
	}
}
