package responses

import (
	modelsGorm "chuck-jokes/pkg/database/models/gorm"
	"time"
)

type UserResponse struct {
	ID        uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Username  string
}

func NewUserResponse(user *modelsGorm.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		Name:      user.Name,
		Username:  user.Username,
	}
}
