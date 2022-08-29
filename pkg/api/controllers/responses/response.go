package responses

import (
	"chuck-jokes/models"
	"chuck-jokes/pkg/repositories"
	"time"
)

type ResponseHandler interface {
	NewUser(user *models.User) User
	NewToken(token string, ttl *time.Time, refreshTTL *time.Time) Token
	NewCategory(category *models.Category) Category
	NewError(err error) Error
	NewErrorsCollection(errors []error) []Error
	NewJoke(joke *models.Joke) Joke
	NewJokeCollection(jokes []models.Joke) []Joke
	NewPaginateJokes(repJokes *repositories.Pagination[models.Joke]) *Pagination[Joke]
	NewSuccess(message string) Success
	NewCategoryJokes(category *models.Category) CategoryJokes
	NewJokeStatistic(joke *models.Joke, favAmount uint) JokeStatistic
}
type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

type Pagination[T interface{}] struct {
	Page       int
	PerPage    int
	TotalRows  int64
	TotalPages int
	Rows       []T
}

func ResponsePagination[T interface{}](
	page int,
	perPage int,
	totalRows int64,
	totalPages int,
	rows []T) *Pagination[T] {
	return &Pagination[T]{Page: page, PerPage: perPage, TotalRows: totalRows, TotalPages: totalPages, Rows: rows}
}
