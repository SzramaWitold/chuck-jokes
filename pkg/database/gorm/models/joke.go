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
	Shows      uint   `gorm:"default:0"`
}

// Create save joke to database
func Create(db *gorm.DB, joke *models.Joke) *models.Joke {
	tx := db.Create(&joke)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	return joke
}
