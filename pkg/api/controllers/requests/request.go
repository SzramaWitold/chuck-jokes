package requests

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestHandler interface {
	NewCreateCategory(c *gin.Context) (*CreateCategory, error)
	NewLogin(c *gin.Context) (*LoginRequest, error)
	NewAddFavouriteRequest(c *gin.Context) (*AddFavourite, error)
	NewFavourites(c *gin.Context) (*Favourites, error)
	NewFindCollection(c *gin.Context) FindCollection
	NewManageCategory(c *gin.Context) (*ManageCategory, error)
	NewJokeOfADay(c *gin.Context) (*JokeOfADay, error)
	NewRegister(c *gin.Context) (*Register, error)
	NewSetAccess(c *gin.Context) (*SetAccess, error)
	NewGetCategory(c *gin.Context) (*GetCategory, error)
	NewJoke(c *gin.Context) (*Joke, error)
}

type RequestValidator struct {
	Validator *validator.Validate
}

func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{
		Validator: validator,
	}
}

func changeToUint(input, name string) (uint, error) {
	i, convErr := strconv.Atoi(input)

	if convErr != nil {
		return 0, fmt.Errorf("field '%v' should be of numeric type", name)
	}

	if i <= 0 {
		return 0, fmt.Errorf("field '%v' should be positive number", name)
	}

	return uint(i), nil
}
