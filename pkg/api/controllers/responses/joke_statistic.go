package responses

import (
	"time"

	"chuck-jokes/models"
)

type JokeStatistic struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Favourite uint
	Shows     uint
	Joke      Joke
}

func (r *DefaultResponseHandler) NewJokeStatistic(js *models.JokeStatistic, favAmount uint) JokeStatistic {
	return JokeStatistic{
		ID:        js.ID,
		CreatedAt: js.CreatedAt,
		UpdatedAt: js.UpdatedAt,
		Shows:     js.Shows,
		Favourite: favAmount,
		Joke:      r.NewJoke(&js.Joke),
	}
}
