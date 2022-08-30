package models

import (
	"time"
)

// User base model.
type User struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string     `faker:"name"`
	Username   string     `faker:"email,unique"`
	Password   string     `faker:"password"`
	Favourites []Joke     `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
