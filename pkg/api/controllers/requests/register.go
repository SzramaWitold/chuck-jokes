package requests

import (
	"github.com/gin-gonic/gin"
)

type Register struct {
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (r *RequestValidator) NewRegister(c *gin.Context) (*Register, error) {
	var request Register

	request.Name = c.PostForm("Name")
	request.Username = c.PostForm("Username")
	request.Password = c.PostForm("Password")

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
