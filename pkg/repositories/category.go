package repositories

import (
	"chuck-jokes/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(userID uint, name string) *models.Category {
	var category = models.Category{
		UserID: userID,
		Name:   name,
	}

	if err := c.db.Model(&category).Where("user_id = ? AND name = ?", userID, name).First(&category).Error; err == nil {
		log.Println("category already exist for user:", userID)
	} else {
		c.db.Create(&category)
	}

	if category.ID == 0 {
		log.Printf("Problem with adding new category with userID: %v and name: %v \n", userID, name)
		return nil
	}

	return &category
}

func (c *Category) AddToCategory(userId, categoryID, jokeID uint) error {
	var category = models.Category{}
	c.db.First(&category, categoryID)
	if category.ID == 0 {
		return fmt.Errorf("can not find category with provided ID: %v", categoryID)
	}
	if category.UserID != userId {
		return fmt.Errorf("do not have permission to add joke to this category")
	}

	var joke = models.Joke{}
	c.db.First(&joke, jokeID)

	if joke.ID == 0 {
		return fmt.Errorf("joke with provided id: %v not exist", jokeID)
	}

	return nil
}

func (c *Category) UpdateAccess(userId, categoryID uint) error {
	var category, categoryError = getCategory(userId, categoryID, c.db)

	if categoryError != nil {
		return categoryError
	}

	category.Access = time.Now().Add(2 * time.Hour)

	tx := c.db.Save(category)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func getCategory(userId, categoryID uint, db *gorm.DB) (*models.Category, error) {
	var category = models.Category{}
	db.First(&category, categoryID)
	if category.ID == 0 {
		return nil, fmt.Errorf("can not find category with provided ID: %v", categoryID)
	}
	if category.UserID != userId {
		return nil, fmt.Errorf("do not have permission to add joke to this category")
	}
	return &category, nil
}
