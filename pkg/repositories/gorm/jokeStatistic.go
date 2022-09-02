package gorm

import (
	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"
	"gorm.io/gorm"
)

type JokeStatisticRepository interface {
	FindByJokeID(jokeID uint) (*models.JokeStatistic, uint, error)
	AddShowByJokeID(jokeID uint) error
	Create(jokeID uint) error
}

type JokeStatistic struct {
	db *gorm.DB
}

func NewJokeStatistic(db *gorm.DB) *JokeStatistic {
	return &JokeStatistic{db: db}
}

func (js *JokeStatistic) Create(jokeID uint) error {
	var jokeStatistic gormModels.JokeStatistic

	jokeStatistic.JokeID = jokeID

	if tx := js.db.Save(&jokeStatistic); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (js *JokeStatistic) AddShowByJokeID(jokeID uint) error {
	var jokeStatistic gormModels.JokeStatistic
	if tx := js.db.Where("joke_id", jokeID).First(&jokeStatistic); tx.Error != nil {
		return tx.Error
	}

	jokeStatistic.Shows = jokeStatistic.Shows + 1

	if saveTX := js.db.Save(jokeStatistic); saveTX.Error != nil {
		return saveTX.Error
	}

	return nil
}

func (js *JokeStatistic) FindByJokeID(jokeID uint) (*models.JokeStatistic, uint, error) {
	var (
		jokeStatistic gormModels.JokeStatistic
		joke          gormModels.Joke
	)

	if queryErr := js.db.Model(&jokeStatistic).Where("joke_id", jokeID).Association("Joke").Append(&joke); queryErr != nil {
		return nil, 0, queryErr
	}

	favourites := js.db.Model(&joke).Association("Users").Count()

	return mapJokeStatistic(jokeStatistic), uint(favourites), nil
}
