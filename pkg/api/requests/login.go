package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username, Password string
}

func NewLogin(c *gin.Context) (*LoginRequest, error) {
	username, password, loginErr := validateLogin(c)
	if loginErr != nil {
		return nil, loginErr
	}
	return &LoginRequest{
		Username: username,
		Password: password,
	}, nil
}

func validateLogin(c *gin.Context) (string, string, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		return "", "", fmt.Errorf("username required")
	}
	if password == "" {
		return "", "", fmt.Errorf("password required")
	}

	return username, password, nil
}
