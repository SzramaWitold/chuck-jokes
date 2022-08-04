package requests

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type FavouritesRequest struct {
	PaginationRequest
	UserID uint
}

func NewFavouritesRequest(c *gin.Context) (*FavouritesRequest, error) {
	userID, userIDErr := validateFavouriteRequest(c)

	if userIDErr != nil {
		return nil, userIDErr
	}

	return &FavouritesRequest{
		PaginationRequest: NewPaginationRequest(c),
		UserID:            userID,
	}, nil
}

func validateFavouriteRequest(c *gin.Context) (uint, error) {
	userID, userIDErr := strconv.Atoi(c.Param("userID"))
	if userIDErr != nil {
		log.Println(userIDErr)
		return 0, userIDErr
	}

	return uint(userID), nil
}
