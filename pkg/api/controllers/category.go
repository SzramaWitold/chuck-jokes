package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cont *Controller) CreateCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewCreateCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		category := cont.Repository.Category.CreateCategory(request.UserID, request.Name)
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

		addError := cont.Repository.Category.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

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
		addError := cont.Repository.Category.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

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

		updateAccessErrors := cont.Repository.Category.UpdateAccess(request.UserID, request.CategoryID)

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

		category, categoryError := cont.Repository.Category.GetCategory(request.UserID, request.CategoryID)

		if categoryError != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(categoryError))
			return
		}

		c.JSON(http.StatusOK, cont.Response.NewCategoryJokes(category))
	}
}
