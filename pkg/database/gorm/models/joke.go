package models

import (
	"gorm.io/gorm"
)

// Joke gorm model needed for migration
type Joke struct {
	gorm.Model
	Value      string `faker:"sentence"`
	ExternalID string `gorm:"unique" faker:"unique"`
	Users      []User `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
