package controllers

import (
	"chuck-jokes/pkg/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (cont *Controller) GetMe() func(c *gin.Context) {
	return func(c *gin.Context) {
		userRepository := repositories.NewUser(cont.DB)
		userID, userIDErr := strconv.Atoi(c.Param("UserID"))
		if userIDErr != nil {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("invalid token, can not fina user")))
			return
		}
		user := userRepository.GetUserFromToken(userID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("can not find user")))
			return
		}
		c.JSON(http.StatusOK, cont.Response.NewUserResponse(user))
	}
}

func (cont *Controller) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		userRepository := repositories.NewUser(cont.DB)
		request, requestErr := cont.Request.NewLogin(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		user := userRepository.Authenticate(request.Username, request.Password)
		if user != nil {
			baseJwt := *cont.JWT
			c.JSON(http.StatusOK, cont.Response.NewTokenResponse(baseJwt.CreateToken(user)))
		} else {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("wrong credentials")))
		}
	}
}

func (cont *Controller) Register() func(c *gin.Context) {
	return func(c *gin.Context) {
		userRepository := repositories.NewUser(cont.DB)
		request, requestErr := cont.Request.NewRegister(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		createUserError := userRepository.Register(request.Name, request.Username, request.Password)

		if createUserError != nil {
			log.Println(createUserError)
			errors := []error{
				fmt.Errorf("something went wrong pleas try again later"),
			}
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(errors))
			return
		}

		c.JSON(http.StatusCreated, cont.Response.NewSuccess("New User created"))
	}
}