package requests

import (
	"github.com/gin-gonic/gin"
)

type AddToCategory struct {
	UserID     uint
	CategoryID uint
	JokeID     uint
}

func (r *Request) NewAddToCategory(c *gin.Context) (*AddToCategory, []error) {
	var errors []error
	errors = append(errors, r.Validator.Validate("userID", c.Param("userID"), "required", "uint")...)
	errors = append(errors, r.Validator.Validate("ID", c.Param("ID"), "required", "uint")...)
	errors = append(errors, r.Validator.Validate("jokeID", c.PostForm("jokeID"), "required", "uint")...)

	if errors != nil {
		return nil, errors
	}
	return &AddToCategory{
		UserID:     changeToUint(c.Param("userID")),
		CategoryID: changeToUint(c.Param("ID")),
		JokeID:     changeToUint(c.PostForm("jokeID")),
	}, nil
}
