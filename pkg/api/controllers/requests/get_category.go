package requests

import (
	"github.com/gin-gonic/gin"
)

type GetCategory struct {
	UserID     uint
	CategoryID uint `validate:"required"`
}

func (r *RequestValidator) NewGetCategory(c *gin.Context) (*GetCategory, error) {
	var request GetCategory

	userID, userIDErr := changeToUint(c.Param("UserID"), "UserID")

	if userIDErr != nil {
		return nil, userIDErr
	} else {
		request.UserID = userID
	}

	categoryID, categoryIDErr := changeToUint(c.Param("ID"), "CategoryID")

	if categoryIDErr != nil {
		return nil, categoryIDErr
	} else {
		request.CategoryID = categoryID
	}

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
