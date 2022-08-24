package controllers

import (
	"chuck-jokes/pkg/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cont *Controller) CreateCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		repository := repositories.NewCategory(cont.DB)
		request, requestErr := cont.Request.NewCreateCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		category := repository.CreateCategory(request.UserID, request.Name)
		c.JSON(http.StatusOK, cont.Response.NewCategory(category))
	}
}

func (cont *Controller) AddToCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewManageCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}
		repository := repositories.NewCategory(cont.DB)
		addError := repository.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

		if addError != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(addError))
			return
		}

		c.JSON(http.StatusOK, cont.Response.NewSuccess("success"))
	}
}

func (cont *Controller) RemoveFromCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewManageCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}
		repository := repositories.NewCategory(cont.DB)
		addError := repository.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

		if addError != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(addError))
			return
		}

		c.JSON(http.StatusOK, cont.Response.NewSuccess("success"))
	}
}

func (cont *Controller) SetAccessCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewSetAccess(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}
		repository := repositories.NewCategory(cont.DB)

		updateAccessErrors := repository.UpdateAccess(request.UserID, request.CategoryID)

		if updateAccessErrors != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(updateAccessErrors))
			return
		}

		c.JSON(http.StatusOK, cont.Response.NewSuccess("success"))
	}
}

func (cont *Controller) GetCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewGetCategory(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}
		repository := repositories.NewCategory(cont.DB)

		category, categoryError := repository.GetCategory(request.UserID, request.CategoryID)

		if categoryError != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(categoryError))
			return
		}

		c.JSON(http.StatusOK, cont.Response.NewCategoryJokes(category))
	}
}
