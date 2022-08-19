package repositories

import (
	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User base user repository
type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Register(name, username, password string) error {
	hashPassword, hashPasswordErr := hashPassword(password)
	if hashPasswordErr != nil {
		return hashPasswordErr
	}
	var user = gormModels.User{
		Username: username,
		Name:     name,
		Password: hashPassword,
	}

	tx := u.db.Create(&user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Authenticate get user based on username and password
func (u *User) Authenticate(username, password string) *models.User {
	var user = models.User{}
	u.db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return nil
	}

	if !checkPasswordHash(password, user.Password) {
		return nil
	}

	return &user
}

func (u *User) GetUserFromToken(id int) *models.User {
	var user = models.User{}
	u.db.Where("ID = ?", id).First(&user)

	if user.ID == 0 {
		return nil
	}

	return &user
}

func (u *User) AddFavourite(userID, jokeID uint) error {
	var joke = models.Joke{}
	var user = models.User{}
	u.db.First(&user, userID)
	u.db.First(&joke, jokeID)

	if user.ID == 0 {
		return fmt.Errorf("user with provided id: %v not exist", userID)
	}

	if joke.ID == 0 {
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
