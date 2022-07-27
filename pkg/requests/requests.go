package requests

import (
	"chuck-jokes/pkg/database/models/gorm"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// CallRandom Call ranodm joke from external api
func CallRandom() gorm.ExternalJoke {
	response, err := http.Get(os.Getenv("EXTERNAL_API") + "jokes/random")
	var joke gorm.ExternalJoke
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
