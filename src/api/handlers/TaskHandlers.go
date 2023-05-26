package handlers

import (
	"time"

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

	if err := context.Bind(&taskRequest); err != nil {

		return err
	}

	currentTime := time.Now()

	taskInstance, err := task.NewBuilder().WithID(uuid.New()).WithTitle(taskRequest.Title).WithCompleted(taskRequest.Completed).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription(taskRequest.Description).Build()

	if err != nil {
		return err
	}

	taskID, _ := instance.service.CreateTask(*taskInstance)

	return context.JSON(200, taskID)

}

func (instance TaskHandlers) GetTask(context echo.Context) error {
	id := uuid.MustParse(context.Param("id"))

	dataTask, err := instance.service.GetTask(id)

	if err != nil {
		return err
	}

	return context.JSON(200, response.Task{
		ID:          dataTask.ID(),
		Title:       dataTask.Title(),
		Description: dataTask.Description(),
		Completed:   dataTask.Completed(),
		CreatedAt:   dataTask.CreatedAt(),
		UpdatedAt:   dataTask.UpdatedAt(),
	})
}

func (instance TaskHandlers) GetTasks(context echo.Context) error {
	listTasks, err := instance.service.GetTasks()

	if err != nil {
		return err
	}

	tasksServices := []response.Task{}

	for _, task := range listTasks {
		tasksServices = append(tasksServices, response.Task{
			ID:          task.ID(),
			Title:       task.Title(),
			Description: task.Description(),
			Completed:   task.Completed(),
			CreatedAt:   task.CreatedAt(),
			UpdatedAt:   task.UpdatedAt(),
		})
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

	taskInstance, err := task.NewBuilder().WithID(id).WithTitle(dataTask.Title).WithCompleted(dataTask.Completed).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription(dataTask.Description).Build()

	if err != nil {
		return err
	}

	data, err := instance.service.UpdateTask(*taskInstance)

	if err != nil {
		return err

	}

	return context.JSON(200, data)

}

func (instance TaskHandlers) DeleteTask(context echo.Context) error {
	id := uuid.MustParse(context.Param("id"))

	err := instance.service.DeleteTask(id)

	if err != nil {
		return err
	}

	return context.JSON(200, nil)

}

func NewTaskHandlers(service primary.TaskManager) *TaskHandlers {
	return &TaskHandlers{service: service}
}
