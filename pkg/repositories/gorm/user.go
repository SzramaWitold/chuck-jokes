package gorm

import (
	"fmt"
	"log"

	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(name string, username string, password string) error
	Authenticate(username string, password string) *models.User
	FindById(id int) *models.User
	AddFavourite(userID uint, jokeID uint) error
}

// User base mapUser repository
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

	user := gormModels.User{
		Username: username,
		Name:     name,
		Password: hashPassword,
	}

	if tx := u.db.Create(&user); tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Authenticate get mapUser based on username and password
func (u *User) Authenticate(username, password string) *models.User {
	user := gormModels.User{}

	if tx := u.db.Where("username = ?", username).First(&user); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	if !checkPasswordHash(password, user.Password) {
		return nil
	}

	return mapUser(&user)
}

func (u *User) FindById(id int) *models.User {
	user := gormModels.User{}

	if tx := u.db.Where("ID = ?", id).First(&user); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	return mapUser(&user)
}

func (u *User) AddFavourite(userID, jokeID uint) error {
	joke := gormModels.Joke{}
	user := gormModels.User{}

	if userTx := u.db.First(&user, userID); userTx.Error != nil {
		log.Println(userTx.Error)

		return userTx.Error
	}

	if jokeTx := u.db.First(&joke, jokeID); jokeTx.Error != nil {
		log.Println(jokeTx.Error)

		return fmt.Errorf("mapJoke with provided id: %v not exist", jokeID)
	}

	if updateError := u.db.Model(&user).Association("Favourites").Append(&joke); updateError != nil {
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
