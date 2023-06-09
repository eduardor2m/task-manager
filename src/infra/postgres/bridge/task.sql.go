// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: task.sql

package bridge

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :exec

INSERT INTO task (id, title, description, category, status, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, title, description, category, status, date, created_at, updated_at
`

type CreateTaskParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	Category    string
	Status      bool
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Category,
		arg.Status,
		arg.Date,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteAllTasks = `-- name: DeleteAllTasks :exec

DELETE FROM task
`

func (q *Queries) DeleteAllTasks(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllTasks)
	return err
}

const deleteTask = `-- name: DeleteTask :exec

DELETE FROM task WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one

SELECT id, title, description, category, status, date, created_at, updated_at FROM task WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Category,
		&i.Status,
		&i.Date,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTasks = `-- name: GetTasks :many

SELECT id, title, description, category, status, date, created_at, updated_at FROM task ORDER BY id DESC
`

func (q *Queries) GetTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Category,
			&i.Status,
			&i.Date,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec

UPDATE task SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5
`

type UpdateTaskParams struct {
	Title       string
	Description string
	Status      bool
	UpdatedAt   time.Time
	ID          uuid.UUID
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
