package gorm

import "gorm.io/gorm"

// User gorm model
type User struct {
	gorm.Model
	Name       string     `faker:"name"`
	Username   string     `faker:"email,unique"`
	Password   string     `faker:"password"`
	Favourites []Joke     `gorm:"many2many:jokes_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categories []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// GetID from User
func (u *User) GetID() uint {
	return u.ID
}
