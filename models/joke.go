package models

import (
	"database/sql"
	"time"
)

type Joke struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
	Shows      uint
	Value      string
	ExternalID string
	Users      []User `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
