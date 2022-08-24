package requests

import (
	"github.com/gin-gonic/gin"
)

type SetAccess struct {
	CategoryID uint `validation:"Required,Uint"`
	UserID     uint `validation:"Required,Uint"`
}

func (r *RequestValidator) NewSetAccess(c *gin.Context) (*SetAccess, []error) {
	inputParams := map[string]string{
		"CategoryID": c.Param("ID"),
		"UserID":     c.Param("UserID"),
	}

	var request SetAccess
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.CategoryID = changeToUint(c.Param("ID"))
	request.UserID = changeToUint(c.Param("UserID"))

	return &request, nil
}
