package database

import (
	"fmt"

	"gorm.io/gorm"
)

// Seeder base struct for initialize seed
type Seeder struct {
	Gorm *gorm.DB
}

// NewSeeder create new gorm seeder
func NewSeeder(gorm *gorm.DB) Seeder {
	return Seeder{
		Gorm: gorm,
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
	user   User
	amount int
	jokes  int
}

// Seed seed database withufake data
func (s *Seeder) Seed() {
	fmt.Println("Seed in progress...")
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
	fmt.Println("Seed complete")
}

// CreateJokes create and return specify amount of jokes
func (s *Seeder) CreateJokes(request JokeCreateRequest) []*Joke {
	jokes := make([]*Joke, request.amount)
	for i := 0; i < request.amount; i++ {
		jokes[i] = CreateJoke(nil)
	}

	return jokes
}

// CreateUsers create with favourite jokes
func (s *Seeder) CreateUsers(request UserCreateRequest) []*User {
	users := make([]*User, request.amount)
	for i := 0; i < request.amount; i++ {
		user := CreateUser(nil)
		for j := 0; j < request.favourites; j++ {
			joke := CreateJoke(nil)
			user.Favourites = append(user.Favourites, *joke)
		}
		users[i] = user
		s.Gorm.Save(&user)
	}

	return users
}

// CreateCategories create categories with attached jokes
func (s *Seeder) CreateCategories(request CategoryRequest) {
	for i := 0; i < request.amount; i++ {
		category := CreateCategory(&request.user, nil)
		for j := 0; j < request.jokes; j++ {
			joke := CreateJoke(nil)
			category.Jokes = append(category.Jokes, *joke)
		}
		s.Gorm.Save(&category)
	}
}
