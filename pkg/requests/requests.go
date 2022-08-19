package requests

import (
	"chuck-jokes/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type IExternalRequest interface {
	CallRandom() *models.Joke
}

type ExternalRequest struct {
	Random string
}

func NewExternalRequest(random string) IExternalRequest {
	return &ExternalRequest{Random: random}
}

// Joke for call from external api
type Joke struct {
	Value string
	ID    string
}

func (j *Joke) ChangeToBaseModel() *models.Joke {
	return &models.Joke{
		Value:      j.Value,
		ExternalID: j.ID,
	}
}

// CallRandom Call random joke from external api
func (e ExternalRequest) CallRandom() *models.Joke {
	response, responseError := http.Get(e.Random)
	var joke Joke
	if responseError != nil {
		log.Println(responseError)
	}

	body, readBodyError := ioutil.ReadAll(response.Body)

	if readBodyError != nil {
		panic(readBodyError)
	}

	jsonUnmarshalError := json.Unmarshal(body, &joke)
	if jsonUnmarshalError != nil {
		panic(jsonUnmarshalError)
	}

	return joke.ChangeToBaseModel()
}
