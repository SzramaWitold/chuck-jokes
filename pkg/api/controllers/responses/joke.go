package responses

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"chuck-jokes/pkg/repositories"
	"time"
)

type Joke struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Value     string
}

func (r *Response) NewJoke(joke *gormModels.Joke) Joke {
	return Joke{
		ID:        joke.ID,
		CreatedAt: joke.CreatedAt,
		UpdatedAt: joke.UpdatedAt,
		Value:     joke.Value,
	}
}

func (r *Response) NewJokeCollection(jokes []gormModels.Joke) []Joke {
	var jokesCollection []Joke

	for _, joke := range jokes {
		jokesCollection = append(jokesCollection, r.NewJoke(&joke))
	}

	return jokesCollection
}

func (r *Response) PaginateJokes(repJokes *repositories.Pagination[gormModels.Joke]) *Pagination[Joke] {
	return ResponsePagination[Joke](repJokes.Page, repJokes.PerPage, repJokes.TotalRows, repJokes.TotalPages, r.NewJokeCollection(repJokes.Rows))
}
