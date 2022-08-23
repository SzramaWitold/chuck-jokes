package requests

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `validation:"Required"`
	Password string `validation:"Required"`
}

func (r *Request) NewLogin(c *gin.Context) (*LoginRequest, []error) {
	inputParams := map[string]string{
		"Username": c.PostForm("Username"),
		"Password": c.PostForm("Password"),
	}
	var request LoginRequest
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.Username = c.PostForm("Username")
	request.Password = c.PostForm("Password")

	return &request, nil
}
