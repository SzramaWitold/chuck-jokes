package requests

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type AddFavouriteRequest struct {
	UserID uint
	JokeID uint
}

func NewAddFavouriteRequest(c *gin.Context) (*AddFavouriteRequest, error) {
	userID, jokeID, err := validateAddFavouriteRequest(c)
	if err != nil {
		return nil, err
	}
	return &AddFavouriteRequest{
		UserID: userID,
		JokeID: jokeID,
	}, nil
}

func validateAddFavouriteRequest(c *gin.Context) (uint, uint, error) {
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
