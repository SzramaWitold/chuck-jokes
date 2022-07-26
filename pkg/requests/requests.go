package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// JokeResponse Base joke response from external api
type JokeResponse struct {
	Value string `faker:"sentence"`
	ID    string `gorm:"primaryKey" faker:"unique"`
}

// CallRandom Call ranodm joke from external api
func CallRandom() JokeResponse {
	response, err := http.Get(os.Getenv("EXTERNAL_API") + "jokes/random")
	var joke JokeResponse
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &joke)

	return joke
}
