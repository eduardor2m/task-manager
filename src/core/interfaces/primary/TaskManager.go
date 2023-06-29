package primary

import (
	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/google/uuid"
)

type TaskManager interface {
	CreateTask(task task.Task) (*uuid.UUID, error)
	GetTask(id uuid.UUID) (*task.Task, error)
	GetTasks() ([]*task.Task, error)
	UpdateTask(task task.Task) (*task.Task, error)
	UpdateTaskStatus(id uuid.UUID) (*task.Task, error)
	DeleteTask(id uuid.UUID) error
	DeleteTasks() error
}
