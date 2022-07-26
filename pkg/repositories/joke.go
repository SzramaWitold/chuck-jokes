package repositories

import "chuck-jokes/pkg/database"

type Joke interface {
	JokeOfTheDay(time string) *database.Joke
	GetJokes(page, perPage int) *Pagination
	JokeExistInLastMonth(joke *database.Joke) bool
}
