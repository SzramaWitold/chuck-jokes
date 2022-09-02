package requests

import (
	"github.com/gin-gonic/gin"
)

type Joke struct {
	JokeID uint `validate:"required"`
}

func (r *RequestValidator) NewJoke(c *gin.Context) (*Joke, error) {
	var request Joke

	jokeID, jokeIDErr := changeToUint(c.Param("ID"))

	if jokeIDErr != nil {
		return nil, jokeIDErr
	}

	request.JokeID = jokeID

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
