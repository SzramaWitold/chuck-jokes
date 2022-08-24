package models

import (
	"time"
)

// Category base model
type Category struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Access    *time.Time
	Name      string
	Jokes     []Joke `gorm:"many2many:categories_jokes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;"`
}
