package gorm

import "gorm.io/gorm"

type Repository struct {
	Category      CategoryRepository
	Joke          JokeRepository
	User          UserRepository
	JokeStatistic JokeStatisticRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category:      NewCategory(db),
		Joke:          NewJoke(db),
		User:          NewUser(db),
		JokeStatistic: NewJokeStatistic(db),
	}
}
