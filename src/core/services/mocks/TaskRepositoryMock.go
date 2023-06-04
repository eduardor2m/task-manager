package mock_repository

import (
	"time"

	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

type MockTaskRepository struct {
	repository.TaskLoader
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{}
}

func (instance MockTaskRepository) CreateTask(task task.Task) (*uuid.UUID, error) {
	return nil, nil
}

func (instance MockTaskRepository) GetTask(id uuid.UUID) (*task.Task, error) {
	taskBuilder := task.NewBuilder()
	time := time.Now()
	id = uuid.New()
	taskDB, _ := taskBuilder.WithID(id).WithTitle("Title").WithDescription("Description").WithCompleted(false).WithCreatedAt(&time).WithUpdatedAt(&time).Build()
	return taskDB, nil
}

func (instance MockTaskRepository) GetTasks() ([]*task.Task, error) {
	return nil, nil
}

func (instance MockTaskRepository) UpdateTask(task task.Task) (*task.Task, error) {
	return nil, nil
}

func (instance MockTaskRepository) DeleteTask(id uuid.UUID) error {
	return nil
}

func (instance MockTaskRepository) DeleteTasks() error {
	return nil
}
