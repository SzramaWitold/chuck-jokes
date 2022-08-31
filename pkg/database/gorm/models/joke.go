package models

import (
	"log"

	"chuck-jokes/models"

	"gorm.io/gorm"
)

// Joke gorm model needed for migration
type Joke struct {
	gorm.Model
	Value      string `faker:"sentence"`
	ExternalID string `gorm:"unique" faker:"unique"`
	Shows      uint   `gorm:"default:0"`
	Users      []User `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Create save joke to database
func Create(db *gorm.DB, joke *models.Joke) *models.Joke {
	if tx := db.Create(&joke); tx.Error != nil {
		log.Println(tx.Error)
	}

	return joke
}
