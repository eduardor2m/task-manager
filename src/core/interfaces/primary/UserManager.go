package primary

import (
	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/google/uuid"
)

type UserManager interface {
	SignUp(user user.User) (*uuid.UUID, error)
	SignIn(username, password string) (*uuid.UUID, error)
}
