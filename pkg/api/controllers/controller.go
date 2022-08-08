package controllers

import (
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/validator"
	"gorm.io/gorm"
)

type Controller struct {
	DB       *gorm.DB
	Request  *requests.Request
	Response *responses.Response
}

func NewController(db *gorm.DB, validator *validator.Validator) *Controller {
	return &Controller{
		DB:       db,
		Request:  requests.NewRequest(validator),
		Response: responses.NewResponse(),
	}
}
