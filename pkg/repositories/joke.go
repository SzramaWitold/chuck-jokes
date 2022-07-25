package repositories

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database"
)

// JokeOfTheDay get joke of the day from specyfi day
func JokeOfTheDay(time string) *database.Joke {
	db := di.GORM()
	var joke = database.Joke{}
	db.Where("created_at >= ?", time).First(&joke)

	if joke.ID == "" {
		return nil
	}
	return &joke
}

// GetJokes get all jokes
func GetJokes() []database.Joke {
	db := di.GORM()
	var jokes = []database.Joke{}
	db.Find(&jokes)

	return jokes
}