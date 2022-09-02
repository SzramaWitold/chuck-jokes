package models

import (
	"time"
)

// User base model.
type User struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Username   string
	Password   string
	Favourites []Joke
	Categories []Category
}
