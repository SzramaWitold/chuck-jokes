package requests

import (
	"github.com/gin-gonic/gin"
)

type Joke struct {
	JokeID uint `validation:"Required,Uint"`
}

func (r *Request) NewJoke(c *gin.Context) (*Joke, []error) {
	inputParams := map[string]string{
		"JokeID": c.Param("ID"),
	}

	var request Joke
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.JokeID = changeToUint(c.Param("ID"))

	return &request, nil
}
