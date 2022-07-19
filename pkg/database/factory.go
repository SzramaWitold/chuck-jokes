package database

import (
	"github.com/bxcodec/faker/v3"
)

// CreateJoke add fake Joke model to database
func CreateJoke(joke *Joke) *Joke {
	if joke == nil {
		joke = new(Joke)
		faker.FakeData(&joke)
	}
	DB.Create(joke)
	return joke
}

// CreateUser add fake User model to database
func CreateUser(user *User) *User {
	if user == nil {
		user = new(User)
		user.Name = faker.Name()
		user.Password = faker.Password()
		user.Username = faker.Email()
	}
	DB.Create(user)
	return user
}

// CreateCategory add fake User category model to database
func CreateCategory(user *User, category *Category) *Category {
	if user == nil {
		panic("User required for category")
	}

	if category == nil {
		category = new(Category)
		category.Name = faker.Name()
		category.UserID = user.ID
	}
	DB.Create(category)

	return category
}


