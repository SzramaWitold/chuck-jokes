package requests

import (
	"chuck-jokes/pkg/validator"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type IRequest interface {
	NewCreateCategory(c *gin.Context) (*CreateCategory, []error)
	NewLogin(c *gin.Context) (*LoginRequest, []error)
	NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, []error)
	NewFavourites(c *gin.Context) (*Favourites, error)
	NewPagination(c *gin.Context) PaginationRequest
	NewAddToCategory(c *gin.Context) (*AddToCategory, []error)
	NewJokeOfADay(c *gin.Context) (*JokeOfADay, []error)
	NewRegister(c *gin.Context) (*Register, []error)
	NewSetAccess(c *gin.Context) (*SetAccess, []error)
}

type Request struct {
	Validator *validator.Validator
}

func NewRequest(validator *validator.Validator) *Request {
	return &Request{
		Validator: validator,
	}
}

func validateTokenUser(c *gin.Context) (uint, error) {
	userID, userIDErr := strconv.Atoi(c.Param("UserID"))
	if userIDErr != nil {
		log.Println(userIDErr)
		return 0, userIDErr
	}

	return uint(userID), nil
}

func changeToUint(input string) uint {
	userID, _ := strconv.Atoi(input) // already validate by Validator

	return uint(userID)
}
