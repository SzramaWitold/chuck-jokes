package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	UserID uint
	Name   string
}

func NewCreateCategory(c *gin.Context) (*CreateCategory, error) {
	userID, name, requestError := validateCreateCategory(c)

	if requestError != nil {
		return nil, requestError
	}
	return &CreateCategory{
		UserID: userID,
		Name:   name,
	}, nil
}

func validateCreateCategory(c *gin.Context) (uint, string, error) {
	userID, userErr := validateTokenUser(c)
	if userErr != nil {
		return 0, "", userErr
	}
	categoryName := c.PostForm("name")

	if categoryName == "" {
		return 0, "", fmt.Errorf("name field required")
	}

	return userID, categoryName, nil
}
