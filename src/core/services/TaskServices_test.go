package services

import (
	"testing"

	mock_repository "github.com/eduardor2m/task-manager/src/core/services/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	mockRepo := mock_repository.NewMockTaskRepository()
	taskServices := NewTaskServices(mockRepo)

	task, err := taskServices.GetTask(uuid.New())

	if err != nil {
		t.Error("Error should be nil")
	}

	assert.Equal(t, task.Title(), "Title")
	assert.Equal(t, task.Description(), "Description")
	assert.Equal(t, task.Status(), false)
}
