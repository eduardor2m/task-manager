package postgres

import (
	"context"
	"github.com/eduardor2m/task-manager/src/infra/postgres/bridge"
	"strconv"
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
	db, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(db)

	ctx := context.Background()

	queries := bridge.New(db)

	err = queries.CreateTask(ctx, bridge.CreateTaskParams{
		Title:       taskInstance.Title(),
		Description: taskInstance.Description(),
		Completed:   taskInstance.Completed(),
		CreatedAt:   taskInstance.CreatedAt(),
		UpdatedAt:   taskInstance.UpdatedAt(),
	})

	lasInsertID := taskInstance.ID()

	return &lasInsertID, nil

}

func (instance TaskSQLiteRepository) GetTask(id uuid.UUID) (*task.Task, error) {
	db, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	smtp, err := db.Prepare("SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer smtp.Close()

	var taskID uuid.UUID
	var taskTitle string
	var taskDescription string
	var taskCompleted string
	var taskCreatedAt time.Time
	var taskUpdatedAt time.Time

	smtp.QueryRow(id).Scan(&taskID, &taskTitle, &taskDescription, &taskCompleted, &taskCreatedAt, &taskUpdatedAt)

	taskCompletedBool, _ := strconv.ParseBool(taskCompleted)

	newTask, _ := task.NewBuilder().WithID(taskID).WithTitle(taskTitle).WithCompleted(taskCompletedBool).WithCreatedAt(&taskCreatedAt).WithUpdatedAt(&taskUpdatedAt).WithDescription(taskDescription).Build()

	return newTask, nil
}

func (instance TaskSQLiteRepository) GetTasks() ([]*task.Task, error) {
	db, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	smtp, err := db.Prepare("SELECT id, title, description, completed, created_at, updated_at FROM tasks")

	if err != nil {
		return nil, err
	}

	defer smtp.Close()

	rows, err := smtp.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*task.Task

	for rows.Next() {
		var taskID uuid.UUID
		var taskTitle string
		var taskDescription string
		var taskCompleted bool
		var taskCreatedAt time.Time
		var taskUpdatedAt time.Time

		_ = rows.Scan(&taskID, &taskTitle, &taskDescription, &taskCompleted, &taskCreatedAt, &taskUpdatedAt)

		newTask, _ := task.NewBuilder().WithID(taskID).WithTitle(taskTitle).WithCompleted(taskCompleted).WithCreatedAt(&taskCreatedAt).WithUpdatedAt(&taskUpdatedAt).WithDescription(taskDescription).Build()

		tasks = append(tasks, newTask)

	}

	return tasks, nil

}

func (instance TaskSQLiteRepository) UpdateTask(taskInstance task.Task) (*task.Task, error) {
	db, err := instance.getConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	smtp, err := db.Prepare("UPDATE tasks SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer smtp.Close()

	taskID := taskInstance.ID()
	taskTitle := taskInstance.Title()
	taskDescription := taskInstance.Description()
	taskCompleted := taskInstance.Completed()
	taskUpdatedAt := taskInstance.UpdatedAt()

	_, err = smtp.Exec(taskTitle, taskDescription, taskCompleted, taskUpdatedAt, taskID)

	if err != nil {
		return nil, err
	}

	return &taskInstance, nil

}

func (instance TaskSQLiteRepository) DeleteTask(id uuid.UUID) error {
	db, err := instance.getConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	smtp, err := db.Prepare("DELETE FROM tasks WHERE id = ?")

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (instance TaskSQLiteRepository) DeleteTasks() error {
	db, err := instance.getConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	smtp, err := db.Prepare("DELETE FROM tasks")

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec()

	if err != nil {
		return err
	}

	return nil
}

func NewTaskSQLiteRepository(manager connectorManager) repository.TaskLoader {
	return &TaskSQLiteRepository{manager}
}
