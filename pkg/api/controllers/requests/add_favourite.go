package requests

import (
	"github.com/gin-gonic/gin"
)

type AddFavourite struct {
	UserID uint `validate:"required"`
	JokeID uint `validate:"required"`
}

func (r *RequestValidator) NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, error) {
	var request AddFavourite

	userID, userIDErr := changeToUint(c.Param("UserID"))

	if userIDErr != nil {
		return nil, userIDErr
	}

	request.UserID = userID

	jokeID, jokeIDErr := changeToUint(c.PostForm("JokeID"))
	if jokeIDErr != nil {
		return nil, jokeIDErr
	} else {
		request.JokeID = jokeID
	}

	reqError := r.Validator.Struct(request)

	if reqError != nil {
		return nil, reqError
	}

	return &request, nil
}
