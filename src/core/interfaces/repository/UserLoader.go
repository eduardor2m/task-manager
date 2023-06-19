package repository

import (
	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/google/uuid"
)

type UserLoader interface {
	SignUp(user user.User) (*uuid.UUID, error)
	SignIn(email string, password string) (*string, error)
}
