package requests

import (
	"github.com/gin-gonic/gin"
)

type SetAccess struct {
	CategoryID uint `validation:"required,uint"`
	UserID     uint `validation:"required,uint"`
}

func (r *Request) NewSetAccess(c *gin.Context) (*SetAccess, []error) {
	inputParams := map[string]string{
		"CategoryID": c.Param("CategoryID"),
		"UserID":     c.Param("CategoryID"),
	}

	var request SetAccess
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.CategoryID = changeToUint(c.Param("CategoryID"))
	request.UserID = changeToUint(c.Param("UserID"))

	return &request, nil
}
