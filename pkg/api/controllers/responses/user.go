package responses

import (
	"chuck-jokes/models"
	"time"
)

type UserResponse struct {
	ID        uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Username  string
}

func (r *Response) NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		Name:      user.Name,
		Username:  user.Username,
	}
}
