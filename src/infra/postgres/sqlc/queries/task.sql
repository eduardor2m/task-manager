-- name: CreateTask :exec

INSERT INTO task (id, title, description, category, status, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetTask :one

SELECT * FROM task WHERE id = $1 LIMIT 1;

-- name: GetTasks :many

SELECT * FROM task ORDER BY id DESC;

-- name: UpdateTask :exec

UPDATE task SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5;

-- name: DeleteTask :exec

DELETE FROM task WHERE id = $1;

-- name: DeleteAllTasks :exec

DELETE FROM task;



