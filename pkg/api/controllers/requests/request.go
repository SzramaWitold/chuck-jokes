package requests

import (
	"chuck-jokes/pkg/validator"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Request struct {
	Validator *validator.Validator
}

func NewRequest(validator *validator.Validator) *Request {
	return &Request{
		Validator: validator,
	}
}

func validateTokenUser(c *gin.Context) (uint, error) {
	userID, userIDErr := strconv.Atoi(c.Param("userID"))
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
