package requests

import (
	"github.com/gin-gonic/gin"
)

type Favourites struct {
	PaginationRequest
	UserID uint
}

func (r *Request) NewFavourites(c *gin.Context) (*Favourites, error) {
	userID, userIDErr := validateTokenUser(c)

	if userIDErr != nil {
		return nil, userIDErr
	}

	return &Favourites{
		PaginationRequest: r.NewPagination(c),
		UserID:            userID,
	}, nil
}
