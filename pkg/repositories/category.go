package repositories

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(userID uint, name string) *gormModels.Category {
	var category = gormModels.Category{
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
