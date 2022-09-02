package requests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"chuck-jokes/models"
)

type ExternalRequestor interface {
	CallRandom() *models.Joke
}

type ExternalRequest struct {
	random string
	client HTTPClient
}

func NewExternalRequest(random string, client HTTPClient) ExternalRequestor {
	return &ExternalRequest{
		random: random,
		client: client,
	}
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
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
	var joke Joke
	request, requestError := http.NewRequest(http.MethodGet, e.random, nil)

	if requestError != nil {
		log.Println(requestError)
	}

	response, responseErr := e.client.Do(request)

	if responseErr != nil {
		log.Println(responseErr)

		return nil
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
