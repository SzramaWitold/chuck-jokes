package repositories

import "gorm.io/gorm"

type Repository struct {
	Category CategoryRepository
	Joke     JokeRepository
	User     UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategory(db),
		Joke:     NewJoke(db),
		User:     NewUser(db),
	}
}
