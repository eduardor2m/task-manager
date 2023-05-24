package response

import (
	"time"

	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func NewTask(taskInstance task.Task) *Task {
	taskResponse := &Task{
		ID:          taskInstance.ID(),
		Title:       taskInstance.Title(),
		Description: taskInstance.Description(),
		Completed:   taskInstance.Completed(),
		CreatedAt:   taskInstance.CreatedAt(),
		UpdatedAt:   taskInstance.UpdatedAt(),
	}

	return taskResponse
}
