package controllers

import (
	"chuck-jokes/pkg/api/requests"
	"chuck-jokes/pkg/api/responses"
	"chuck-jokes/pkg/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cont *Controller) CreateCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		repository := repositories.NewCategory(cont.DB)
		request, requestErr := requests.NewCreateCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, requestErr.Error())
			return
		}

		category := repository.CreateCategory(request.UserID, request.Name)

		c.JSON(http.StatusOK, responses.NewCategory(category))
	}
}

func (cont *Controller) AddToCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, requestErr := requests.NewCreateCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, requestErr.Error())
			return
		}
		c.JSON(http.StatusOK, "I am in")
	}
}
