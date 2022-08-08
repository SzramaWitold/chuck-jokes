package requests

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username, Password string
}

func (r *Request) NewLogin(c *gin.Context) (*LoginRequest, []error) {
	var errors []error
	errors = append(errors, r.Validator.Validate("username", c.PostForm("username"), "required")...)
	errors = append(errors, r.Validator.Validate("password", c.PostForm("password"), "required")...)

	if errors != nil {
		return nil, errors
	}

	return &LoginRequest{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}, nil
}
