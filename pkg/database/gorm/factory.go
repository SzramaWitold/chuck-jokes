package gorm

import (
	"log"
	"math/rand"
	"time"

	"chuck-jokes/models"
	gormRepository "chuck-jokes/pkg/repositories/gorm"

	"github.com/bxcodec/faker/v3"
)

// Factory base struct for create data inside database
type Factory struct {
	repository *gormRepository.Repository
}

// NewFactory create new Factory
func NewFactory(repository *gormRepository.Repository) *Factory {
	return &Factory{
		repository: repository,
	}
}

// CreateJoke add Joke model to database or populate it with fake data
func (f *Factory) CreateJoke(addDays int) *models.Joke {
	rand.Seed(time.Now().Unix())
	joke := models.Joke{
		Value:      faker.Sentence(),
		ExternalID: faker.UUIDHyphenated(),
	}
	joke.CreatedAt = time.Now().Add(24 * time.Duration(addDays) * time.Hour)
	newJoke, createErr := f.repository.Joke.Create(&joke)

	if createErr != nil {
		log.Println(createErr)

		return nil
	}

	return newJoke
}

// CreateUser add fake User model to database
func (f *Factory) CreateUser() *models.User {
	user, createErr := f.repository.User.Create(faker.Name(), faker.Email(), faker.Password())
	if createErr != nil {
		log.Println(createErr)

		return nil
	}

	return user
}

// CreateCategory add fake User category model to database
func (f *Factory) CreateCategory(user *models.User, category *models.Category) *models.Category {
	if user == nil {
		panic("User required for category")
	}

	category, createErr := f.repository.Category.Create(user.ID, faker.Name())

	if createErr != nil {
		log.Println(createErr)

		return nil
	}

	return category
}
