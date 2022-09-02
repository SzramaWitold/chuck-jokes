package responses

import (
	"chuck-jokes/models"
	"time"
)

type User struct {
	ID        uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Username  string
}

func (r *DefaultResponseHandler) NewUser(user *models.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		Name:      user.Name,
		Username:  user.Username,
	}
}
