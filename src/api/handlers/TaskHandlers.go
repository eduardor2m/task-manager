package handlers

import (
	"github.com/eduardor2m/task-manager/src/api/handlers/dto/request"
	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/primary"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskHandlers struct {
	service primary.TaskManager
}

func (instance TaskHandlers) CreateTask(context echo.Context) error {
	var dto request.TaskDTO

	bindError := context.Bind(&dto)

	if bindError != nil {
		return bindError
	}

	taskBuilder := task.NewBuilder()

	taskBuilder.WithID(
		uuid.New(),
	).WithTitle(dto.Title).WithDescription(dto.Description).WithCompleted(false)

	taskInstance, validationError := taskBuilder.Build()

	if validationError != nil {
		return validationError
	}

	task, createError := instance.service.CreateTask(*taskInstance)

	if createError != nil {
		return createError
	}

	return context.JSON(201, task)

}

func (instance TaskHandlers) GetTask(context echo.Context) error {
	return nil
}

func (instance TaskHandlers) GetTasks(context echo.Context) error {
	return context.JSON(200, "Hello World")
}

func (instance TaskHandlers) UpdateTask(context echo.Context) error {
	return nil
}

func (instance TaskHandlers) DeleteTask(context echo.Context) error {
	return nil
}

func NewTaskHandlers(service primary.TaskManager) *TaskHandlers {
	return &TaskHandlers{service: service}
}
