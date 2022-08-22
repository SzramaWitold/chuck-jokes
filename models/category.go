package models

import (
	"database/sql"
	"time"
)

// Category base model
type Category struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Access    time.Time
	Name      string
	Jokes     []Joke `gorm:"many2many:categories_jokes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;"`
}
