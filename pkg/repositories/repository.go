package repositories

import "gorm.io/gorm"

type Repository struct {
	Category ICategory
	Joke     IJoke
	User     IUser
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategory(db),
		Joke:     NewJoke(db),
		User:     NewUser(db),
	}
}
