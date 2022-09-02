package requests

import (
	"github.com/gin-gonic/gin"
)

type SetAccess struct {
	CategoryID uint `validate:"required"`
	UserID     uint `validate:"required"`
}

func (r *RequestValidator) NewSetAccess(c *gin.Context) (*SetAccess, error) {
	var request SetAccess

	categoryID, categoryIDErr := changeToUint(c.Param("ID"))

	if categoryIDErr != nil {
		return nil, categoryIDErr
	} else {
		request.CategoryID = categoryID
	}

	userID, userIDErr := changeToUint(c.Param("UserID"))

	if userIDErr != nil {
		return nil, userIDErr
	} else {
		request.UserID = userID
	}

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
