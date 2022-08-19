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
	Value      string
	ExternalID string
}
