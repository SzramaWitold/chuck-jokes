package responses

import (
	"chuck-jokes/models"
	"time"
)

type JokeStatistic struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Favourite uint
	Shows     uint
}

func (r *Response) NewJokeStatistic(joke *models.Joke, favAmount uint) JokeStatistic {
	return JokeStatistic{
		ID:        joke.ID,
		CreatedAt: joke.CreatedAt,
		UpdatedAt: joke.UpdatedAt,
		Shows:     joke.Shows,
		Favourite: favAmount,
	}
}
