package repositories

import (
	"fmt"
	"log"
	"time"

	"chuck-jokes/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(userID uint, name string) *models.Category
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

func (c *Category) Create(userID uint, name string) *models.Category {
	category := models.Category{
		UserID: userID,
		Name:   name,
	}

	if tx := c.db.Create(&category); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	return &category
}

func (c *Category) AddToCategory(userId, categoryID, jokeID uint) error {
	category, categoryError := getUserCategory(userId, categoryID, c.db)

	if categoryError != nil {
		return categoryError
	}

	joke := models.Joke{}

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
	category := models.Category{}

	if tx := c.db.Preload("Jokes").First(&category, categoryID); tx.Error != nil {
		return nil, tx.Error
	}

	if time.Now().Before(*category.Access) {
		return &category, nil
	}

	if userId != 0 {
		return &category, nil
	}

	return nil, fmt.Errorf("do not have permission to get this category")
}

func getUserCategory(userId, categoryID uint, db *gorm.DB) (*models.Category, error) {
	category := models.Category{}

	if tx := db.First(&category, categoryID); tx.Error != nil {
		return nil, tx.Error
	}

	if category.UserID != userId {
		return nil, fmt.Errorf("do not have permission to add joke to this category")
	}

	return &category, nil
}
