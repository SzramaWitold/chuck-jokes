package responses

import (
	"chuck-jokes/models"
	"log"
	"time"
)

type CategoryJokes struct {
	ID        uint
	CreatedAt *time.Time
	Name      string
	UserID    uint
	Jokes     []Joke
}

func (r *Response) NewCategoryJokes(category *models.Category) CategoryJokes {
	log.Println(category)
	return CategoryJokes{
		ID:        category.ID,
		CreatedAt: &category.CreatedAt,
		Name:      category.Name,
		UserID:    category.UserID,
		Jokes:     r.NewJokeCollection(category.Jokes),
	}
}
