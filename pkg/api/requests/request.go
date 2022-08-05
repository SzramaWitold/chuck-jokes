package requests

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func validateTokenUser(c *gin.Context) (uint, error) {
	userID, userIDErr := strconv.Atoi(c.Param("userID"))
	if userIDErr != nil {
		log.Println(userIDErr)
		return 0, userIDErr
	}

	return uint(userID), nil
}
