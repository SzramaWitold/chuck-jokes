package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JokeResponse struct {
	Value string
	Id    string 
}

func CallRandom() JokeResponse {
	response, err := http.Get("https://api.chucknorris.io/jokes/random")
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
