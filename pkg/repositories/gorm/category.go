package gorm

import (
	"fmt"
	"log"
	"time"

	"chuck-jokes/models"
	gormModel "chuck-jokes/pkg/database/gorm/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(userID uint, name string) (*models.Category, error)
	AddToCategory(userId uint, categoryID uint, jokeID uint) error
	UpdateAccess(userId uint, categoryID uint) error
	FindByUserIDAndCategoryID(userId uint, categoryID uint) (*models.Category, error)
}

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(userID uint, name string) (*models.Category, error) {
	cat := gormModel.Category{
		UserID: userID,
		Name:   name,
	}

	if tx := c.db.Create(&cat); tx.Error != nil {
		log.Println(tx.Error)

		return nil, tx.Error
	}

	return mapCategory(&cat), nil
}

func (c *Category) AddToCategory(userId, categoryID, jokeID uint) error {
	category, categoryError := getUserCategory(userId, categoryID, c.db)

	if categoryError != nil {
		return categoryError
	}

	joke := gormModel.Joke{}

	if tx := c.db.First(&joke, jokeID); tx.Error != nil {
		return tx.Error
	}

	updateError := c.db.Model(&category).Association("Jokes").Append(&joke)

	if updateError != nil {
		return updateError
	}

	return nil
}

func (c *Category) UpdateAccess(userId, categoryID uint) error {
	category, categoryError := getUserCategory(userId, categoryID, c.db)

	if categoryError != nil {
		return categoryError
	}

	access := time.Now().Add(2 * time.Hour)
	category.Access = &access

	if tx := c.db.Save(category); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *Category) FindByUserIDAndCategoryID(userId, categoryID uint) (*models.Category, error) {
	cat := gormModel.Category{}

	if tx := c.db.Preload("Jokes").First(&cat, categoryID); tx.Error != nil {
		return nil, tx.Error
	}

	if time.Now().Before(*cat.Access) {
		return mapCategory(&cat), nil
	}

	if userId != 0 {
		return mapCategory(&cat), nil
	}

	return nil, fmt.Errorf("do not have permission to get this cat")
}

func getUserCategory(userId, categoryID uint, db *gorm.DB) (*gormModel.Category, error) {
	category := gormModel.Category{}

	if tx := db.First(&category, categoryID); tx.Error != nil {
		return nil, tx.Error
	}

	if category.UserID != userId {
		return nil, fmt.Errorf("do not have permission to add mapJoke to this mapCategory")
	}

	return &category, nil
}
