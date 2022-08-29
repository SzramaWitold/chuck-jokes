package requests

import (
	"github.com/gin-gonic/gin"
)

type ManageCategory struct {
	UserID     uint `validate:"required"`
	CategoryID uint `validate:"required"`
	JokeID     uint `validate:"required"`
}

func (r *RequestValidator) NewManageCategory(c *gin.Context) (*ManageCategory, error) {

	request := ManageCategory{}

	userID, userIDErr := changeToUint(c.Param("UserID"), "UserID")

	if userIDErr != nil {
		return nil, userIDErr
	} else {
		request.UserID = userID
	}

	jokeID, jokeIDErr := changeToUint(c.PostForm("JokeID"), "JokeID")
	if jokeIDErr != nil {
		return nil, jokeIDErr
	} else {
		request.JokeID = jokeID
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
