package models

import (
	"time"
)

type Joke struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Shows      uint
	Value      string
	ExternalID string
	Users      []User
}
