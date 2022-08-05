package responses

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"time"
)

type Category struct {
	ID        uint
	CreatedAt *time.Time
	Name      string
	UserID    uint
}

func NewCategory(category *gormModels.Category) *Category {
	return &Category{
		ID:        category.ID,
		CreatedAt: &category.CreatedAt,
		Name:      category.Name,
		UserID:    category.UserID,
	}
}
