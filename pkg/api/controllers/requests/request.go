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
	NewManageCategory(c *gin.Context) (*ManageCategory, []error)
	NewJokeOfADay(c *gin.Context) (*JokeOfADay, []error)
	NewRegister(c *gin.Context) (*Register, []error)
	NewSetAccess(c *gin.Context) (*SetAccess, []error)
	NewGetCategory(c *gin.Context) (*GetCategory, []error)
	NewJoke(c *gin.Context) (*Joke, []error)
}

type RequestValidator struct {
	Validator validator.IValidator
}

func NewRequestValidator(validator validator.IValidator) *RequestValidator {
	return &RequestValidator{
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
