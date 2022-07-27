package gorm

import (
	"log"

	"gorm.io/gorm"
)

// Joke gorm model
type Joke struct {
	gorm.Model
	Value      string `faker:"sentence"`
	ExternalID string `gorm:"unique" faker:"unique"`
}

// Create save joke to database
func (j *Joke) Create(db *gorm.DB) *Joke {
 tx := db.Create(j)
 if tx.Error != nil {
 	log.Println(tx.Error)
 }

 return j
}

// GetID from Joke
func (j *Joke) GetID() uint {
	return j.ID
}


// ExternalJoke for call from external api
type ExternalJoke struct {
	Value string `faker:"sentence"`
	ID    string `gorm:"primaryKey" faker:"unique"`
}

// ChangeToGormJoke transfer data from ExternalJoke to DBJoke
func ChangeToGormJoke(joke *ExternalJoke) Joke {
	return Joke{
		Value:      joke.Value,
		ExternalID: joke.ID,
	}
}
