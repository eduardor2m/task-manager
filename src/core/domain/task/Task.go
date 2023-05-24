package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id          uuid.UUID
	title       string
	description string
	completed   bool
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

func (instance *Task) Completed() bool {
	return instance.completed
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
