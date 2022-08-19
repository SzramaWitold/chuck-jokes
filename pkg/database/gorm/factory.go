package gorm

import (
	"chuck-jokes/pkg/database/gorm/models"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

// Factory base struct for create data inside database
type Factory struct {
	db *gorm.DB
}

// NewFactory create new Factory
func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		db: db,
	}
}

// CreateJoke add Joke model to database or populate it with fake data
func (f *Factory) CreateJoke() *models.Joke {
	joke := models.Joke{
		Value:      faker.Sentence(),
		ExternalID: faker.UUIDHyphenated(),
	}

	f.db.Create(&joke)

	return &joke
}

// CreateUser add fake User model to database
func (f *Factory) CreateUser() *models.User {
	user := models.User{}
	user.Name = faker.Name()
	user.Password = faker.Password()
	user.Username = faker.Email()

	f.db.Create(&user)

	return &user
}

// CreateCategory add fake User category model to database
func (f *Factory) CreateCategory(user *models.User, category *models.Category) *models.Category {
	if user == nil {
		panic("User required for category")
	}

	if category == nil {
		category = new(models.Category)
		category.Name = faker.Name()
		category.UserID = user.ID
	}

	f.db.Create(category)

	return category
}
