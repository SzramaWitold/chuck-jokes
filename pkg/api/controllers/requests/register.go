package requests

import "github.com/gin-gonic/gin"

type Register struct {
	Name     string `validation:"Required"`
	Username string `validation:"Required,Unique:users"`
	Password string `validation:"Required"`
}

func (r *Request) NewRegister(c *gin.Context) (*Register, []error) {
	inputParams := map[string]string{
		"Name":     c.PostForm("Name"),
		"Username": c.PostForm("Username"),
		"Password": c.PostForm("Password"),
	}

	var request Register
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.Name = c.PostForm("Name")
	request.Username = c.PostForm("Username")
	request.Password = c.PostForm("Password")

	return &request, nil
}
