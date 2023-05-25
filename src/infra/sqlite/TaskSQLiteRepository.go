package sqlite

import (
	"fmt"
	"time"

	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ repository.TaskLoader = &TaskSQLiteRepository{}

type TaskSQLiteRepository struct {
	connectorManager
}

func (instance TaskSQLiteRepository) CreateTask(taskInstance task.Task) (*uuid.UUID, error) {
	fmt.Println("CreateTask Infra")
	fmt.Println(taskInstance)

	taskID := taskInstance.ID()
	taskTitle := taskInstance.Title()
	taskDescription := taskInstance.Description()
	taskCompleted := taskInstance.Completed()
	taskCreatedAt := taskInstance.CreatedAt()
	taskUpdatedAt := taskInstance.UpdatedAt()

	fmt.Println(taskID)
	fmt.Println(taskTitle)
	fmt.Println(taskDescription)
	fmt.Println(taskCompleted)
	fmt.Println(taskCreatedAt)
	fmt.Println(taskUpdatedAt)

	return &taskID, nil
}

func (instance TaskSQLiteRepository) GetTask(id uuid.UUID) (*task.Task, error) {
	return nil, nil
}

func (instance TaskSQLiteRepository) GetTasks() ([]*task.Task, error) {
	listTasks := []*task.Task{}

	currentTime := time.Now()
	newTask, _ := task.NewBuilder().WithID(uuid.New()).WithTitle("title").WithCompleted(false).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).WithDescription("description").Build()

	listTasks = append(listTasks, newTask)

	return listTasks, nil
}

func (instance TaskSQLiteRepository) UpdateTask(taskInstance task.Task) (*task.Task, error) {
	return nil, nil
}

func (instance TaskSQLiteRepository) DeleteTask(id uuid.UUID) error {
	return nil
}

func NewTaskSQLiteRepository(manager connectorManager) repository.TaskLoader {
	return &TaskSQLiteRepository{manager}
}
