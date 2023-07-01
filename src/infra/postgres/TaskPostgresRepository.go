package postgres

import (
	"context"
	"time"

	"github.com/eduardor2m/task-manager/src/infra/postgres/bridge"

	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ repository.TaskLoader = &TaskSQLiteRepository{}

type TaskSQLiteRepository struct {
	connectorManager
}

func (instance TaskSQLiteRepository) CreateTask(taskInstance task.Task) (*uuid.UUID, error) {
	db, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(db)

	ctx := context.Background()

	queries := bridge.New(db)

	taskFormated := bridge.CreateTaskParams{
		ID:          taskInstance.ID(),
		Title:       taskInstance.Title(),
		Category:    taskInstance.Category(),
		Description: taskInstance.Description(),
		Status:      taskInstance.Status(),
		Date:        *taskInstance.Date(),
		CreatedAt:   *taskInstance.CreatedAt(),
		UpdatedAt:   *taskInstance.UpdatedAt(),
	}

	err = queries.CreateTask(ctx, taskFormated)

	if err != nil {
		return nil, err
	}

	idLastInsert := taskFormated.ID

	return &idLastInsert, nil

}

func (instance TaskSQLiteRepository) GetTask(id uuid.UUID) (*task.Task, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	taskFormated, err := queries.GetTask(ctx, id)

	if err != nil {
		return nil, err
	}

	taskInstance, _ := task.NewBuilder().WithID(taskFormated.ID).WithTitle(taskFormated.Title).WithCategory(taskFormated.Category).WithStatus(taskFormated.Status).WithCreatedAt(&taskFormated.CreatedAt).WithDate(&taskFormated.Date).WithUpdatedAt(&taskFormated.UpdatedAt).WithDescription(taskFormated.Description).Build()

	return taskInstance, nil
}

func (instance TaskSQLiteRepository) GetTasks() ([]*task.Task, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	tasksFormated, err := queries.GetTasks(ctx)

	if err != nil {
		return nil, err
	}

	var tasks []*task.Task

	for _, taskFormated := range tasksFormated {
		taskInstance, _ := task.NewBuilder().WithID(taskFormated.ID).WithDate(&taskFormated.Date).WithCategory(taskFormated.Category).WithTitle(taskFormated.Title).WithStatus(taskFormated.Status).WithCreatedAt(&taskFormated.CreatedAt).WithUpdatedAt(&taskFormated.UpdatedAt).WithDescription(taskFormated.Description).Build()
		tasks = append(tasks, taskInstance)
	}

	return tasks, nil

}

func (instance TaskSQLiteRepository) UpdateTask(taskInstance task.Task) (*task.Task, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	taskFormated := bridge.UpdateTaskParams{
		ID:          taskInstance.ID(),
		Title:       taskInstance.Title(),
		Description: taskInstance.Description(),
		Status:      taskInstance.Status(),
		UpdatedAt:   *taskInstance.UpdatedAt(),
	}

	updatedTask, err := queries.UpdateTask(ctx, taskFormated)

	if err != nil {
		return nil, err
	}

	formattedTask, _ := task.NewBuilder().WithID(updatedTask.ID).WithTitle(updatedTask.Title).WithStatus(updatedTask.Status).WithUpdatedAt(&updatedTask.UpdatedAt).WithDescription(updatedTask.Description).WithCategory(updatedTask.Category).WithCreatedAt(&updatedTask.CreatedAt).WithDate(&updatedTask.Date).Build()

	return formattedTask, nil

}

func (instance TaskSQLiteRepository) UpdateTaskStatus(id uuid.UUID) (*task.Task, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	taskForUpdate, err := instance.GetTask(id)

	if err != nil {
		return nil, err
	}

	var newStatus bool
	oldStatus := taskForUpdate.Status()

	if oldStatus {
		newStatus = false
	} else {
		newStatus = true
	}

	dateUpdated := time.Now()

	taskFormated := bridge.UpdateTaskStatusParams{
		ID:        id,
		Status:    newStatus,
		UpdatedAt: dateUpdated,
	}

	updatedTask, err := queries.UpdateTaskStatus(ctx, taskFormated)

	if err != nil {
		return nil, err
	}

	formattedTask, err := task.NewBuilder().WithID(updatedTask.ID).WithTitle(updatedTask.Title).WithStatus(updatedTask.Status).WithUpdatedAt(&updatedTask.UpdatedAt).WithDescription(updatedTask.Description).WithCategory(updatedTask.Category).WithCreatedAt(&updatedTask.CreatedAt).WithDate(&updatedTask.Date).Build()

	if err != nil {
		return nil, err
	}

	return formattedTask, nil

}

func (instance TaskSQLiteRepository) DeleteTask(id uuid.UUID) error {
	conn, err := instance.getConnection()

	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.DeleteTask(ctx, id)

	if err != nil {
		return err
	}

	return nil

}

func (instance TaskSQLiteRepository) DeleteTasks() error {
	conn, err := instance.getConnection()

	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.DeleteAllTasks(ctx)

	if err != nil {
		return err
	}

	return nil
}

func NewTaskSQLiteRepository(manager connectorManager) repository.TaskLoader {
	return &TaskSQLiteRepository{manager}
}
