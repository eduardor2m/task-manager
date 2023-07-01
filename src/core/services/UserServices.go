package services

import (
	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

type UserServices struct {
	taskRepository repository.UserLoader
}

func (instance UserServices) SignUp(user user.User) (*uuid.UUID, error) {
	return instance.taskRepository.SignUp(user)
}

func (instance UserServices) SignIn(email string, password string) (*string, error) {
	return instance.taskRepository.SignIn(email, password)
}

func (instance UserServices) DeleteUserByEmail(email string) error {
	return instance.taskRepository.DeleteUserByEmail(email)
}

func NewUserServices(taskRepository repository.UserLoader) UserServices {
	return UserServices{taskRepository: taskRepository}
}
