package requests

import (
	"github.com/gin-gonic/gin"
)

type AddToCategory struct {
	UserID     uint `validation:"required,uint"`
	CategoryID uint `validation:"required,uint"`
	JokeID     uint `validation:"required,uint"`
}

func (r *Request) NewAddToCategory(c *gin.Context) (*AddToCategory, []error) {
	inputParams := map[string]string{
		"UserID": c.Param("UserID"),
		"JokeID": c.PostForm("JokeID"),
		"ID":     c.Param("ID"),
	}
	request := AddToCategory{}
	errors := r.Validator.Validate(&request, inputParams)
	if errors != nil {
		return nil, errors
	}

	request.UserID = changeToUint(c.Param("UserID"))
	request.JokeID = changeToUint(c.PostForm("JokeID"))
	request.CategoryID = changeToUint(c.Param("ID"))

	return &request, nil
}
