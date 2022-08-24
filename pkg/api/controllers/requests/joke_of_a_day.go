package requests

import (
	"github.com/gin-gonic/gin"
)

type JokeOfADay struct {
	Date string `validation:"Date"`
}

func (r *RequestValidator) NewJokeOfADay(c *gin.Context) (*JokeOfADay, []error) {
	inputParams := map[string]string{
		"Date": c.Query("date"),
	}

	var request JokeOfADay
	errors := r.Validator.Validate(request, inputParams)

	if errors != nil {
		return nil, errors
	}

	request.Date = c.Query("date")

	return &request, nil
}
