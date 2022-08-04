package repositories

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository base user repository
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Authenticate get user based on username and password
func (u *UserRepository) Authenticate(username, password string) *gormModels.User {
	var user = gormModels.User{}
	u.db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return nil
	}

	if !checkPasswordHash(password, user.Password) {
		return nil
	}

	return &user
}

func (u *UserRepository) AddFavourite(userID, jokeID uint) error {
	var joke = gormModels.Joke{}
	var user = gormModels.User{}
	u.db.First(&user, userID)
	u.db.First(&joke, jokeID)

	if user.ID == 0 {
		return fmt.Errorf("user with provided id: %v not exist", userID)
	}

	if user.ID == 0 {
		return fmt.Errorf("joke with provided id: %v not exist", jokeID)
	}

	updateError := u.db.Model(&user).Association("Favourites").Append(&joke)

	if updateError != nil {
		return updateError
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
