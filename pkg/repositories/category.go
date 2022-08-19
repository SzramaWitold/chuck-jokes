package repositories

import (
	"chuck-jokes/models"
	"fmt"
	"gorm.io/gorm"
	"log"
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
