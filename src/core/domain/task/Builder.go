package task

import (
	"errors"
	"time"

	"github.com/eduardor2m/task-manager/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	task *Task
	err  error
}

func NewBuilder() *builder {
	return &builder{
		task: &Task{},
	}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.err = errors.New("id is not valid")
		return instance
	}
	instance.task.id = id
	return instance
}

func (instance *builder) WithTitle(title string) *builder {
	if !validator.IsTitleValid(title) {
		instance.err = errors.New("title is not valid")
		return instance
	}

	instance.task.title = title
	return instance
}

func (instance *builder) WithDescription(description string) *builder {
	if !validator.IsDescriptionValid(description) {
		instance.err = errors.New("description is not valid")
		return instance
	}

	instance.task.description = description
	return instance
}

func (instance *builder) WithCompleted(completed bool) *builder {
	instance.task.completed = completed
	return instance
}

func (instance *builder) WithCreatedAt(createdAt *time.Time) *builder {
	// if !validator.IsDateValid(*createdAt) {
	// 	instance.err = errors.New("created at is not valid")
	// 	return instance
	// }

	instance.task.createdAt = createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	// if !validator.IsDateValid(*updatedAt) {
	// 	instance.err = errors.New("updated at is not valid")
	// 	return instance
	// }

	instance.task.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Task, error) {
	if instance.err != nil {
		return nil, instance.err
	}

	return instance.task, nil
}
