package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id          uuid.UUID
	title       string
	description string
	status      bool
	category    string
	date        *time.Time
	createdAt   *time.Time
	updatedAt   *time.Time
}

func (instance *Task) ID() uuid.UUID {
	return instance.id
}

func (instance *Task) Title() string {
	return instance.title
}

func (instance *Task) Description() string {
	return instance.description
}

func (instance *Task) Category() string {
	return instance.category
}

func (instance *Task) Status() bool {
	return instance.status
}

func (instance *Task) Date() *time.Time {
	return instance.date
}

func (instance *Task) CreatedAt() *time.Time {
	return instance.createdAt
}

func (instance *Task) UpdatedAt() *time.Time {
	return instance.updatedAt
}

func (instance *Task) IsZero() bool {
	return instance == &Task{}
}
