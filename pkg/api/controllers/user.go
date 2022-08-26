package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetMe godoc
// @Summary      GetMe
// @Description  get information about user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object}  responses.User
// @Failure      401  {object}  responses.Error
// @Router       /me [get]
func (cont *Controller) GetMe() func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, userIDErr := strconv.Atoi(c.Param("UserID"))
		if userIDErr != nil {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("invalid token, can not fina user")))
			return
		}
		user := cont.Repository.User.GetUserFromToken(userID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("can not find user")))
			return
		}
		c.JSON(http.StatusOK, cont.Response.NewUser(user))
	}
}

// Login godoc
// @Summary      Login
// @Description  get JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Username    formData     string  true  "Username"
// @Param        Password    formData     string  true  "Password"
// @Success      200  {object}   responses.Token
// @Failure      401  {object}  responses.Error
// @Router       /login [post]
func (cont *Controller) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewLogin(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		user := cont.Repository.User.Authenticate(request.Username, request.Password)
		if user != nil {
			baseJwt := *cont.JWT
			c.JSON(http.StatusOK, cont.Response.NewToken(baseJwt.CreateToken(user)))
		} else {
			c.JSON(http.StatusUnauthorized, cont.Response.NewError(fmt.Errorf("wrong credentials")))
		}
	}
}

// Register godoc
// @Summary      Register
// @Description  create new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Username    formData     string  true  "Username"
// @Param        Name    formData     string  true  "Name"
// @Param        Password    formData     string  true  "Password"
// @Success      200  {object}   responses.Success
// @Failure      400  {array}  responses.Error
// @Router       /register [post]
func (cont *Controller) Register() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewRegister(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		}

		createUserError := cont.Repository.User.Register(request.Name, request.Username, request.Password)

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

// GetFavourites godoc
// @Summary      GetFavourites
// @Description  get user favourites jokes
// @Tags         User
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.PaginateJokes
// @Failure      401  {object}  responses.Error
// @Router       /favourite [get]
func (cont *Controller) GetFavourites() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewFavourites(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewError(requestErr))
		}
		jokes := cont.Repository.Joke.GetFavourites(request.Page, request.PerPage, request.UserID)

		c.JSON(http.StatusOK, cont.Response.NewPaginateJokes(jokes))
	}
}

// AddFavourite godoc
// @Summary      AddFavourite
// @Description  Add new favourite joke
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        JokeID    formData     number  true  "JokeID"
// @Param Authorization header string true "With the bearer started"
// @Success      200  {object} responses.Success
// @Failure      400  {array}  responses.Error
// @Failure      417  {object}  responses.Error
// @Router       /favourite [put]
func (cont *Controller) AddFavourite() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := cont.Request.NewAddFavouriteRequest(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, cont.Response.NewErrorsCollection(requestErr))
			return
		} else {
			repErr := cont.Repository.User.AddFavourite(request.UserID, request.JokeID)

			if repErr != nil {
				c.JSON(http.StatusExpectationFailed, cont.Response.NewError(repErr))
				return
			}
			c.JSON(http.StatusOK, cont.Response.NewSuccess("success"))
		}
	}
}
