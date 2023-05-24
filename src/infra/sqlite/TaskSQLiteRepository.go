package sqlite

import (
	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ repository.TaskLoader = &TaskSQLiteRepository{}

type TaskSQLiteRepository struct {
	connectorManager
}

func (instance TaskSQLiteRepository) CreateTask(taskInstance task.Task) (*uuid.UUID, error) {
	conn, err := instance.connectorManager.getConnection()

	if err != nil {
		return nil, err
	}

	defer instance.connectorManager.closeConnection(conn)

	stmt, err := conn.Prepare("INSERT INTO tasks(id, title, description, completed, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(taskInstance.ID(), taskInstance.Title(), taskInstance.Description(), taskInstance.Completed(), taskInstance.CreatedAt(), taskInstance.UpdatedAt())

	if err != nil {
		return nil, err
	}

	taskID := taskInstance.ID()

	return &taskID, nil

}

func (instance TaskSQLiteRepository) GetTask(id uuid.UUID) (*task.Task, error) {
	return nil, nil
}

func (instance TaskSQLiteRepository) GetTasks() ([]*task.Task, error) {
	return nil, nil
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
