package requests

import (
	"github.com/gin-gonic/gin"
)

type Favourites struct {
	PaginationRequest
	UserID uint
}

func NewFavourites(c *gin.Context) (*Favourites, error) {
	userID, userIDErr := validateTokenUser(c)

	if userIDErr != nil {
		return nil, userIDErr
	}

	return &Favourites{
		PaginationRequest: NewPagination(c),
		UserID:            userID,
	}, nil
}
