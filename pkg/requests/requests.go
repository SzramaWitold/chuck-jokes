package requests

import (
	"chuck-jokes/pkg/database/models/gorm"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// CallRandom Call random joke from external api
func CallRandom() gorm.ExternalJoke {
	response, responseError := http.Get(os.Getenv("EXTERNAL_API") + "jokes/random")
	var joke gorm.ExternalJoke
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

	return joke
}
