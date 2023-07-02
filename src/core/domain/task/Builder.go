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

func (instance *builder) WithUserID(userID uuid.UUID) *builder {
	if !validator.IsUUIDValid(userID) {
		instance.err = errors.New("user id is not valid")
		return instance
	}
	instance.task.userID = userID
	return instance
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

func (instance *builder) WithCategory(category string) *builder {
	instance.task.category = category
	return instance
}

func (instance *builder) WithStatus(completed bool) *builder {
	instance.task.status = completed
	return instance
}

func (instance *builder) WithDate(date *time.Time) *builder {
	instance.task.date = date
	return instance
}

func (instance *builder) WithCreatedAt(createdAt *time.Time) *builder {
	if !validator.IsDateValid(*createdAt) {
		instance.err = errors.New("created at is not valid")
		return instance
	}

	instance.task.createdAt = createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	if !validator.IsDateValid(*updatedAt) {
		instance.err = errors.New("updated at is not valid")
		return instance
	}

	instance.task.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Task, error) {
	if instance.err != nil {
		return nil, instance.err
	}

	return instance.task, nil
}
