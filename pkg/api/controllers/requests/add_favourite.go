package requests

import (
	"github.com/gin-gonic/gin"
)

type AddFavourite struct {
	UserID uint `validation:"Required,Uint"`
	JokeID uint `validation:"Required,Uint"`
}

func (r *RequestValidator) NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, []error) {
	inputParams := map[string]string{
		"UserID": c.Param("UserID"),
		"JokeID": c.PostForm("JokeID"),
	}

	var request AddFavourite
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.UserID = changeToUint(c.Param("UserID"))
	request.JokeID = changeToUint(c.PostForm("JokeID"))

	return &request, nil
}
