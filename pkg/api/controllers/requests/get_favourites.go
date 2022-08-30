package requests

import (
	"github.com/gin-gonic/gin"
)

type Favourites struct {
	FindCollection
	UserID uint `validate:"required"`
}

func (r *RequestValidator) NewFavourites(c *gin.Context) (*Favourites, error) {
	var request Favourites

	userID, userIDErr := changeToUint(c.Param("UserID"), "UserID")

	if userIDErr != nil {
		return nil, userIDErr
	}

	request.UserID = userID

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
