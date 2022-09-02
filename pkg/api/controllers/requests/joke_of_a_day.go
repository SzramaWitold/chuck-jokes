package requests

import (
	"github.com/gin-gonic/gin"
)

type JokeOfADay struct {
	Date string `validate:"datetime"`
}

func (r *RequestValidator) NewJokeOfADay(c *gin.Context) (*JokeOfADay, error) {
	var request JokeOfADay

	request.Date = c.Query("date")

	requestError := r.Validator.Struct(request)

	if requestError != nil {
		return nil, requestError
	}

	return &request, nil
}
