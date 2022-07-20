package database

import (
	"chuck-jokes/pkg/requests"
	"time"

	"gorm.io/gorm"
)

// Joke gorm model
type Joke struct {
	requests.JokeResponse
	CreatedAt time.Time
}

// User gorm model
type User struct {
	gorm.Model
	Name       string     `faker:"name"`
	Username   string     `faker:"email,unique"`
	Password   string     `faker:"password"`
	Favourites []Joke     `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Category gorm model
type Category struct {
	gorm.Model
	Name   string
	Jokes  []Joke `gorm:"many2many:categories_jokes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID uint
	User   User `gorm:"foreignKey:UserID;"`
}

// GetAllModels for gorm migrations
func GetAllModels() []interface{} {
	var models = make([]interface{}, 3)
	models[0] = &Joke{}
	models[1] = &User{}
	models[2] = &Category{}

	return models
}
