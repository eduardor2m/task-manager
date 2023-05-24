package handlers

import (
	"fmt"
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

	fmt.Println(taskRequest)

	currentTime := time.Now()

	taskInstance, err := task.NewBuilder().WithID(uuid.New()).WithTitle("title").WithCompleted(false).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription("description").Build()

	if err != nil {
		return err
	}

	taskID, _ := instance.service.CreateTask(*taskInstance)

	return context.JSON(200, taskID)

}

func (instance TaskHandlers) GetTask(context echo.Context) error {
	return nil
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
	return nil
}

func (instance TaskHandlers) DeleteTask(context echo.Context) error {
	return nil
}

func NewTaskHandlers(service primary.TaskManager) *TaskHandlers {
	return &TaskHandlers{service: service}
}
