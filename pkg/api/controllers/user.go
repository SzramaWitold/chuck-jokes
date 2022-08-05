package controllers

import (
	"chuck-jokes/pkg/api/requests"
	"chuck-jokes/pkg/api/responses"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (cont *Controller) GetMe() func(c *gin.Context) {
	return func(c *gin.Context) {
		userRepository := repositories.NewUser(cont.DB)
		userID, userIDErr := strconv.Atoi(c.Param("userID"))
		if userIDErr != nil {
			c.JSON(http.StatusUnauthorized, "Invalid token, can not fina user")
			return
		}
		user := userRepository.GetUserFromToken(userID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, "Can not find user")
			return
		}
		c.JSON(http.StatusOK, responses.NewUserResponse(user))
	}
}

func (cont *Controller) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		userRepository := repositories.NewUser(cont.DB)
		request, requestErr := requests.NewLogin(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, requestErr.Error())
			return
		}

		user := userRepository.Authenticate(request.Username, request.Password)
		if user != nil {
			validator := token.NewHandler()

			c.JSON(http.StatusOK, responses.NewTokenResponse(validator.CreateToken(user)))
		} else {
			c.JSON(http.StatusUnauthorized, "Wrong credentials!")
		}
	}
}
