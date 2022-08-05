package requests

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddToCategory struct {
	UserID     uint
	CategoryID uint
	JokeID     uint
}

func NewAddToCategory(c *gin.Context) (*AddToCategory, error) {
	userID, categoryID, jokeID, requestError := validateAddToCategory(c)
	if requestError != nil {
		return nil, requestError
	}
	return &AddToCategory{
		UserID:     userID,
		CategoryID: categoryID,
		JokeID:     jokeID,
	}, nil
}

func validateAddToCategory(c *gin.Context) (uint, uint, uint, error) {
	userID, userErr := validateTokenUser(c)
	if userErr != nil {
		return 0, 0, 0, userErr
	}
	categoryID, catErr := strconv.Atoi(c.Param("ID"))

	if catErr != nil {
		return 0, 0, 0, catErr
	}

	jokeID, jokeErr := strconv.Atoi(c.PostForm("jokeID"))

	if jokeErr != nil {
		return 0, 0, 0, catErr
	}

	return userID, uint(categoryID), uint(jokeID), nil
}
