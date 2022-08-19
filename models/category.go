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
	Name      string
	Jokes     []Joke
	UserID    uint
	User      User
}
