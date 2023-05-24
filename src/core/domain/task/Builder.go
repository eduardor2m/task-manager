package task

import (
	"time"

	"github.com/google/uuid"
)

type builder struct {
	task *Task
}

func NewBuilder() *builder {
	return &builder{
		task: &Task{},
	}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	instance.task.id = id
	return instance
}

func (instance *builder) WithTitle(title string) *builder {
	instance.task.title = title
	return instance
}

func (instance *builder) WithDescription(description string) *builder {
	instance.task.description = description
	return instance
}

func (instance *builder) WithCompleted(completed bool) *builder {
	instance.task.completed = completed
	return instance
}

func (instance *builder) WithCreatedAt(createdAt *time.Time) *builder {
	instance.task.createdAt = createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	instance.task.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Task, error) {
	return instance.task, nil
}
