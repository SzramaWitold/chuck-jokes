package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	UserID uint   `validate:"required"`
	Name   string `validate:"required"`
}

func (r *RequestValidator) NewCreateCategory(c *gin.Context) (*CreateCategory, error) {
	var request CreateCategory

	userID, userIDErr := changeToUint(c.Param("UserID"))

	if userIDErr != nil {
		return nil, userIDErr
	} else {
		request.UserID = userID
	}
	request.Name = c.PostForm("Name")

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
