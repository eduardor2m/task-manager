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
	Category    string     `json:"category"`
	Status      bool       `json:"status"`
	Date        *time.Time `json:"date"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type TaskID struct {
	ID uuid.UUID `json:"id"`
}

type TaskMessage struct {
	Message string `json:"message"`
}

func NewTask(taskInstance task.Task) *Task {
	taskResponse := &Task{
		ID:          taskInstance.ID(),
		Title:       taskInstance.Title(),
		Description: taskInstance.Description(),
		Category:    taskInstance.Category(),
		Status:      taskInstance.Status(),
		Date:        taskInstance.Date(),
		CreatedAt:   taskInstance.CreatedAt(),
		UpdatedAt:   taskInstance.UpdatedAt(),
	}

	return taskResponse
}
