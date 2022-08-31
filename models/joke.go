package models

import (
	"time"
)

type Joke struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Value      string
	ExternalID string
	Users      []User
}
