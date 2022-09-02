package controllers

import (
	"net/http"

	"chuck-jokes/pkg/repositories/gorm"

	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory() func(c *gin.Context)
	AddToCategory() func(c *gin.Context)
	SetAccessCategory() func(c *gin.Context)
	RemoveFromCategory() func(c *gin.Context)
	GetCategory() func(c *gin.Context)
}

type Category struct {
	request    requests.RequestHandler
	response   responses.ResponseHandler
	repository *gorm.Repository
}

func NewCategory(request requests.RequestHandler, response responses.ResponseHandler, repository *gorm.Repository) CategoryController {
	return &Category{request: request, response: response, repository: repository}
}

// CreateCategory godoc
// @Summary      CreateCategory
// @Description  Create new user category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        Name    formData     string  true  "Name"
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Category
// @Failure      400  {object}  responses.Error
// @Router       /categories [post]
func (cat *Category) CreateCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cat.request.NewCreateCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(requestErr))

			return
		}

		category, databaseErr := cat.repository.Category.Create(request.UserID, request.Name)

		if databaseErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(databaseErr))

			return
		}

		c.JSON(http.StatusOK, cat.response.NewCategory(category))
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
// @Router       /categories/{ID} [put]
func (cat *Category) AddToCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cat.request.NewManageCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(requestErr))
			return
		}

		addError := cat.repository.Category.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

		if addError != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(addError))
			return
		}

		c.JSON(http.StatusOK, cat.response.NewSuccess("success"))
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
// @Router       /categories/{ID} [delete]
func (cat *Category) RemoveFromCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cat.request.NewManageCategory(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(requestErr))
			return
		}
		addError := cat.repository.Category.AddToCategory(request.UserID, request.CategoryID, request.JokeID)

		if addError != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(addError))
			return
		}

		c.JSON(http.StatusOK, cat.response.NewSuccess("success"))
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
// @Router       /categories/{ID}/access [put]
func (cat *Category) SetAccessCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cat.request.NewSetAccess(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(requestErr))
			return
		}

		updateAccessErrors := cat.repository.Category.UpdateAccess(request.UserID, request.CategoryID)

		if updateAccessErrors != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(updateAccessErrors))
			return
		}

		c.JSON(http.StatusOK, cat.response.NewSuccess("success"))
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
// @Router       /categories/{ID} [get]
func (cat *Category) GetCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cat.request.NewGetCategory(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(requestErr))
			return
		}

		category, categoryError := cat.repository.Category.FindByUserIDAndCategoryID(request.UserID, request.CategoryID)

		if categoryError != nil {
			c.JSON(http.StatusBadRequest, cat.response.NewError(categoryError))
			return
		}

		c.JSON(http.StatusOK, cat.response.NewCategoryJokes(category))
	}
}
