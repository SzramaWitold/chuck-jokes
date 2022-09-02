package models

import (
	"time"
)

// Category base model.
type Category struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Access    *time.Time
	Name      string
	Jokes     []Joke
	UserID    uint
	User      User
}
