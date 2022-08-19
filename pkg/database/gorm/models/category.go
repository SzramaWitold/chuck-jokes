package models

import "gorm.io/gorm"

// Category gorm model
type Category struct {
	gorm.Model
	Name   string
	Jokes  []Joke `gorm:"many2many:categories_jokes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID uint
	User   User `gorm:"foreignKey:UserID;"`
}

// GetID from Category
func (c *Category) GetID() uint {
	return c.ID
}
