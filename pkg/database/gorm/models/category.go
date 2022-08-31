package models

import (
	"time"

	"gorm.io/gorm"
)

// Category gorm model.
type Category struct {
	gorm.Model
	Name   string `gorm:"index:idx_owner,unique"`
	Jokes  []Joke `gorm:"many2many:categories_jokes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Access *time.Time
	UserID uint `gorm:"index:idx_owner,unique"`
	User   User `gorm:"foreignKey:UserID;"`
}
