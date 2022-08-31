package models

import "gorm.io/gorm"

type JokeStatistic struct {
	gorm.Model
	Shows  uint `gorm:"default:0"`
	JokeID uint
	Joke   Joke `gorm:"foreignKey:JokeID;"`
}
