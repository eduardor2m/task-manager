package handlers

import (
	"net/http"
	"time"

	"github.com/eduardor2m/task-manager/src/api/handlers/dto"
	"github.com/eduardor2m/task-manager/src/api/handlers/dto/request"
	"github.com/eduardor2m/task-manager/src/api/handlers/dto/response"
	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/primary"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskHandlers struct {
	service primary.TaskManager
}

func (instance TaskHandlers) CreateTask(context echo.Context) error {
	var taskRequest request.TaskDTO

	err := context.Bind(&taskRequest)

	if err != nil {
		return err
	}

	currentTime := time.Now()

	taskInstance, err := task.NewBuilder().WithID(uuid.New()).WithTitle(taskRequest.Title).WithDate(&currentTime).WithCategory(taskRequest.Category).WithStatus(taskRequest.Status).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription(taskRequest.Description).Build()

	if err != nil {
		errMessage := dto.ErrorMessage{
			Message: err.Error(),
		}

		return context.JSON(http.StatusBadRequest, errMessage)
	}

	taskID, _ := instance.service.CreateTask(*taskInstance)

	return context.JSON(200, taskID)
}

func (instance TaskHandlers) GetTask(context echo.Context) error {
	id := uuid.MustParse(context.Param("id"))

	responseTask, err := instance.service.GetTask(id)

	if err != nil {
		return err
	}

	formattedTask := response.NewTask(*responseTask)

	return context.JSON(200, formattedTask)
}

func (instance TaskHandlers) GetTasks(context echo.Context) error {
	listTasks, err := instance.service.GetTasks()

	if err != nil {
		return err
	}

	tasksServices := []response.Task{}

	for _, task := range listTasks {
		tasksServices = append(tasksServices, *response.NewTask(*task))
	}

	return context.JSON(200, tasksServices)
}

func (instance TaskHandlers) UpdateTask(context echo.Context) error {
	dataTask := request.TaskDTO{}

	if err := context.Bind(&dataTask); err != nil {
		return err
	}

	id := uuid.MustParse(context.Param("id"))

	currentTime := time.Now()

	taskInstance, err := task.NewBuilder().WithID(id).WithTitle(dataTask.Title).WithCategory(dataTask.Category).WithStatus(dataTask.Status).WithDate(&currentTime).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription(dataTask.Description).Build()

	if err != nil {
		return err
	}

	data, err := instance.service.UpdateTask(*taskInstance)

	if err != nil {
		return err

	}

	dataFormatted := response.Task{
		ID:          data.ID(),
		Title:       data.Title(),
		Description: data.Description(),
		Category:    data.Category(),
		Status:      data.Status(),
		Date:        data.Date(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
	}

	return context.JSON(200, dataFormatted)

}

func (instance TaskHandlers) DeleteTask(context echo.Context) error {
	id := uuid.MustParse(context.Param("id"))

	err := instance.service.DeleteTask(id)

	if err != nil {
		return err
	}

	message := "Task deleted successfully"

	return context.JSON(200, message)

}

func (instance TaskHandlers) DeleteTasks(context echo.Context) error {
	err := instance.service.DeleteTasks()

	if err != nil {
		return err
	}

	message := "Tasks deleted successfully"

	return context.JSON(200, message)
}

func NewTaskHandlers(service primary.TaskManager) *TaskHandlers {
	return &TaskHandlers{service: service}
}
