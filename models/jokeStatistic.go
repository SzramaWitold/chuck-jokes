package models

import "time"

type JokeStatistic struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Shows     uint
	JokeID    uint
	Joke      Joke
}
