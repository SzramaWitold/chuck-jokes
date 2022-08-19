package requests

import (
	"github.com/gin-gonic/gin"
)

type JokeOfADay struct {
	Date string `validation:"date"`
}

func (r *Request) NewJokeOfADay(c *gin.Context) (*JokeOfADay, []error) {
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
