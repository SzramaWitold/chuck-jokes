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
	Users      []User `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
