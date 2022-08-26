package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCategory godoc
// @Summary      CreateCategory
// @Description  Create new user category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        Name    formData     string  true  "Name"
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Category
// @Failure      400  {array}  responses.Error
// @Router       /categories [post]
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

// AddToCategory godoc
// @Summary      AddToCategory
// @Description  Add joke to category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        JokeID    formData     number  true  "JokeID"
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Success
// @Failure      400  {object}  responses.Error
// @Failure      400  {array}  responses.Error
// @Router       /categories/{ID} [put]
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

// RemoveFromCategory godoc
// @Summary      RemoveFromCategory
// @Description  Remove joke to category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        JokeID    formData     number  true  "JokeID"
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Success
// @Failure      400  {object}  responses.Error
// @Failure      400  {array}  responses.Error
// @Router       /categories/{ID} [delete]
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

// SetAccessCategory godoc
// @Summary      SetAccessCategory
// @Description  Set time limit access for guest
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Success
// @Failure      400  {object}  responses.Error
// @Failure      400  {array}  responses.Error
// @Router       /categories/{ID}/access [put]
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

// GetCategory godoc
// @Summary      GetCategory
// @Description  get category with jokes if access is set then for everyone otherwise for logged user only
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        Name    formData     string  true  "Name"
// @Param Authorization header string false "With the bearer started"
// @Success      200  {object} responses.Category
// @Failure      400  {object}  responses.Error
// @Failure      400  {array}  responses.Error
// @Router       /categories/{ID} [get]
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
