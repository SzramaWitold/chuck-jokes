package requests

import (
	"github.com/gin-gonic/gin"
)

type GetCategory struct {
	UserID     uint
	CategoryID uint `validation:"Required,Uint"`
}

func (r *Request) NewGetCategory(c *gin.Context) (*GetCategory, []error) {
	inputParams := map[string]string{
		"UserID":     c.Param("UserID"),
		"CategoryID": c.Param("ID"),
	}

	var request GetCategory
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.UserID = changeToUint(c.Param("UserID"))
	request.CategoryID = changeToUint(c.Param("ID"))

	return &request, nil
}
