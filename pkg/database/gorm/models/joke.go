package models

import (
	"chuck-jokes/models"
	"log"

	"gorm.io/gorm"
)

// Joke gorm model needed for migration
type Joke struct {
	gorm.Model
	Value      string `faker:"sentence"`
	ExternalID string `gorm:"unique" faker:"unique"`
}

// Create save joke to database
func Create(db *gorm.DB, joke *models.Joke) *models.Joke {
	tx := db.Create(&joke)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	return joke
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
