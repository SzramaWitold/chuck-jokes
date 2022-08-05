package requests

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type AddFavourite struct {
	UserID uint
	JokeID uint
}

func NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, error) {
	userID, jokeID, err := validateAddFavourite(c)
	if err != nil {
		return nil, err
	}
	return &AddFavourite{
		UserID: userID,
		JokeID: jokeID,
	}, nil
}

func validateAddFavourite(c *gin.Context) (uint, uint, error) {
	userID, userIDErr := strconv.Atoi(c.Param("userID"))
	jokeID, jokeErr := strconv.Atoi(c.PostForm("jokeID"))
	if userIDErr != nil {
		log.Println(userIDErr)
		return 0, 0, userIDErr
	}
	if jokeErr != nil {
		log.Println(jokeErr)
		return 0, 0, jokeErr
	}

	return uint(userID), uint(jokeID), nil
}
