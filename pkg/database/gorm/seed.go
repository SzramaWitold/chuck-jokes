package gorm

import (
	"chuck-jokes/pkg/database/gorm/models"
	"log"

	"gorm.io/gorm"
)

// Seeder base struct for initialize seed
type Seeder struct {
	factory *Factory
}

// NewSeeder create new gorm seeder
func NewSeeder(gorm *gorm.DB) Seeder {
	return Seeder{
		factory: NewFactory(gorm),
	}
}

// UserCreateRequest request for seed users
type UserCreateRequest struct {
	amount     int
	favourites int
}

// JokeCreateRequest request for seed users
type JokeCreateRequest struct {
	amount int
}

// CategoryRequest request for add user categories with jokes
type CategoryRequest struct {
	user   models.User
	amount int
	jokes  int
}

// Seed database with fake data
func (s *Seeder) Seed() {
	log.Println("Seed in progress...")
	userRequest := UserCreateRequest{
		amount:     5,
		favourites: 5,
	}
	jokeRequest := JokeCreateRequest{
		amount: 10,
	}
	s.CreateJokes(jokeRequest)
	users := s.CreateUsers(userRequest)
	for _, user := range users {
		categoryRequest := CategoryRequest{
			user:   *user,
			amount: 3,
			jokes:  5,
		}
		s.CreateCategories(categoryRequest)
	}
	log.Println("Seed complete")
}

// CreateJokes create and return specify amount of jokes
func (s *Seeder) CreateJokes(request JokeCreateRequest) []*models.Joke {
	jokes := make([]*models.Joke, request.amount)
	for i := 0; i < request.amount; i++ {
		jokes[i] = s.factory.CreateJoke(-1 * i)
	}

	return jokes
}

// CreateUsers create with favourite jokes
func (s *Seeder) CreateUsers(request UserCreateRequest) []*models.User {
	users := make([]*models.User, request.amount)
	for i := 0; i < request.amount; i++ {
		user := s.factory.CreateUser()
		for j := 0; j < request.favourites; j++ {
			joke := s.factory.CreateJoke(-1 * j)
			user.Favourites = append(user.Favourites, *joke)
		}
		users[i] = user
		s.factory.db.Save(&user)
	}

	return users
}

// CreateCategories create categories with attached jokes
func (s *Seeder) CreateCategories(request CategoryRequest) {
	for i := 0; i < request.amount; i++ {
		category := s.factory.CreateCategory(&request.user, nil)
		for j := 0; j < request.jokes; j++ {
			joke := s.factory.CreateJoke(-1 * j)
			category.Jokes = append(category.Jokes, *joke)
		}
		s.factory.db.Save(&category)
	}
}
