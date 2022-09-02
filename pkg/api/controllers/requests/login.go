package requests

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (r *RequestValidator) NewLogin(c *gin.Context) (*LoginRequest, error) {
	var request LoginRequest

	request.Username = c.PostForm("Username")
	request.Password = c.PostForm("Password")

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
