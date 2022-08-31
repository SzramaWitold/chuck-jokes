package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"chuck-jokes/pkg/repositories/gorm"

	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/token"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetMe() func(c *gin.Context)
	Login() func(c *gin.Context)
	GetFavourites() func(c *gin.Context)
	AddFavourite() func(c *gin.Context)
	Register() func(c *gin.Context)
}

type User struct {
	request    requests.RequestHandler
	response   responses.ResponseHandler
	jwt        *token.TokenHandler
	repository *gorm.Repository
}

func NewUser(request requests.RequestHandler, response responses.ResponseHandler, repository *gorm.Repository, JWT *token.TokenHandler) *User {
	return &User{request: request, response: response, jwt: JWT, repository: repository}
}

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
func (u *User) GetMe() func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, userIDErr := strconv.Atoi(c.Param("UserID"))
		if userIDErr != nil {
			c.JSON(http.StatusUnauthorized, u.response.NewError(fmt.Errorf("invalid token, can not fina user")))

			return
		}

		user := u.repository.User.FindById(userID)

		if user == nil {
			c.JSON(http.StatusUnauthorized, u.response.NewError(fmt.Errorf("can not find user")))

			return
		}

		c.JSON(http.StatusOK, u.response.NewUser(user))
	}
}

// Login godoc
// @Summary      Login
// @Description  get jwt token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Username    formData     string  true  "Username"
// @Param        Password    formData     string  true  "Password"
// @Success      200  {object}   responses.Token
// @Failure      401  {object}  responses.Error
// @Router       /login [post]
func (u *User) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := u.request.NewLogin(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(requestErr))

			return
		}

		user, userErr := u.repository.User.Get(request.Username)

		if userErr != nil || user == nil {
			c.JSON(http.StatusUnauthorized, u.response.NewError(userErr))

			return
		}

		if !checkPasswordHash(request.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, u.response.NewError(fmt.Errorf("wrong credentials")))

			return
		}

		baseJwt := *u.jwt
		c.JSON(http.StatusOK, u.response.NewToken(baseJwt.CreateToken(user)))
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
func (u *User) Register() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := u.request.NewRegister(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(requestErr))

			return
		}

		hashPassword, hashPasswordErr := hashPassword(request.Password)

		if hashPasswordErr != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(hashPasswordErr))

			return
		}

		_, createUserError := u.repository.User.Create(request.Name, request.Username, hashPassword)

		if createUserError != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(createUserError))

			return
		}

		c.JSON(http.StatusCreated, u.response.NewSuccess("New User created"))
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
func (u *User) GetFavourites() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := u.request.NewFavourites(c)
		if requestErr != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(requestErr))

			return
		}

		jokes := u.repository.Joke.FindFavourites(request.FindCollection, request.UserID)

		c.JSON(http.StatusOK, u.response.NewPaginateJokes(jokes))
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
func (u *User) AddFavourite() func(c *gin.Context) {
	return func(c *gin.Context) {
		request, requestErr := u.request.NewAddFavouriteRequest(c)

		if requestErr != nil {
			c.JSON(http.StatusBadRequest, u.response.NewError(requestErr))

			return
		}

		repErr := u.repository.User.AddFavourite(request.UserID, request.JokeID)

		if repErr != nil {
			c.JSON(http.StatusExpectationFailed, u.response.NewError(repErr))

			return
		}

		c.JSON(http.StatusOK, u.response.NewSuccess("success"))
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
