package user

import (
	"errors"
	"time"

	"github.com/eduardor2m/task-manager/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	user *User
	err  error
}

func NewBuilder() *builder {
	return &builder{
		user: &User{},
	}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.err = errors.New("id is not valid")
		return instance
	}
	instance.user.id = id
	return instance
}

func (instance *builder) WithUsername(username string) *builder {
	instance.user.username = username
	return instance
}

func (instance *builder) WithEmail(email string) *builder {
	instance.user.email = email
	return instance
}

func (instance *builder) WithPassword(password string) *builder {
	instance.user.password = password
	return instance
}

func (instance *builder) WithCreatedAt(createdAt *time.Time) *builder {
	if !validator.IsDateValid(*createdAt) {
		instance.err = errors.New("created at is not valid")
		return instance
	}

	instance.user.createdAt = createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	if !validator.IsDateValid(*updatedAt) {
		instance.err = errors.New("updated at is not valid")
		return instance
	}

	instance.user.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*User, error) {
	if instance.err != nil {
		return nil, instance.err
	}

	return instance.user, nil
}
