package requests

import (
	"github.com/gin-gonic/gin"
)

type AddFavourite struct {
	UserID uint
	JokeID uint
}

func (r *Request) NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, []error) {
	var errors []error
	errors = append(errors, r.Validator.Validate("userID", c.Param("userID"), "required", "uint")...)
	errors = append(errors, r.Validator.Validate("jokeID", c.PostForm("jokeID"), "required", "uint")...)

	if errors != nil {
		return nil, errors
	}

	return &AddFavourite{
		UserID: changeToUint(c.Param("userID")),
		JokeID: changeToUint(c.PostForm("jokeID")),
	}, nil
}
