package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	username  string
	email     string
	password  string
	createdAt *time.Time
	updatedAt *time.Time
}

func (instance *User) ID() uuid.UUID {
	return instance.id
}

func (instance *User) Username() string {
	return instance.username
}

func (instance *User) Email() string {
	return instance.email
}

func (instance *User) Password() string {
	return instance.password
}

func (instance *User) CreatedAt() *time.Time {
	return instance.createdAt
}

func (instance *User) UpdatedAt() *time.Time {
	return instance.updatedAt
}

func (instance *User) IsZero() bool {
	return instance == &User{}
}
