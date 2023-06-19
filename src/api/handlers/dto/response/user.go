package response

import (
	"time"

	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewUser(userInstance user.User) *User {
	userResponse := &User{
		ID:        userInstance.ID(),
		Username:  userInstance.Username(),
		Email:     userInstance.Email(),
		Password:  userInstance.Password(),
		CreatedAt: userInstance.CreatedAt(),
		UpdatedAt: userInstance.UpdatedAt(),
	}

	return userResponse
}
