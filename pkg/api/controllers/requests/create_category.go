package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	UserID uint
	Name   string
}

func (r *Request) NewCreateCategory(c *gin.Context) (*CreateCategory, []error) {
	var errors []error
	errors = append(errors, r.Validator.Validate("userID", c.Param("userID"), "required", "uint")...)
	errors = append(errors, r.Validator.Validate("name", c.PostForm("name"), "required")...)

	if errors != nil {
		return nil, errors
	}
	return &CreateCategory{
		UserID: changeToUint(c.Param("userID")),
		Name:   c.PostForm("name"),
	}, nil
}
