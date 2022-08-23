package requests

import (
	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	UserID uint   `validation:"Required,Uint"`
	Name   string `validation:"Required"`
}

func (r *Request) NewCreateCategory(c *gin.Context) (*CreateCategory, []error) {

	inputParams := map[string]string{
		"UserID": c.Param("UserID"),
		"Name":   c.PostForm("Name"),
	}
	var request CreateCategory
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.UserID = changeToUint(c.Param("UserID"))
	request.Name = c.PostForm("Name")

	return &request, nil
}
