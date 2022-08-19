package responses

import (
	"chuck-jokes/models"
	"time"
)

type Category struct {
	ID        uint
	CreatedAt *time.Time
	Name      string
	UserID    uint
}

func (r *Response) NewCategory(category *models.Category) Category {
	return Category{
		ID:        category.ID,
		CreatedAt: &category.CreatedAt,
		Name:      category.Name,
		UserID:    category.UserID,
	}
}
