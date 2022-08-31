package gorm

import (
	"fmt"
	"log"

	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(name string, username string, password string) (*models.User, error)
	Get(username string) (*models.User, error)
	FindById(id int) *models.User
	AddFavourite(userID uint, jokeID uint) error
}

// User base user repository
type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Create(name, username, password string) (*models.User, error) {
	user := gormModels.User{
		Username: username,
		Name:     name,
		Password: password,
	}

	if tx := u.db.Create(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return mapUser(&user), nil
}

// Get user based on username and password
func (u *User) Get(username string) (*models.User, error) {
	user := gormModels.User{}

	if tx := u.db.Where("username = ?", username).First(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return mapUser(&user), nil
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
