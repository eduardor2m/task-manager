-- name: CreateTask :exec

INSERT INTO task (title, description, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5 );

-- name: GetTask :one

SELECT * FROM task WHERE id = $1 LIMIT 1;

-- name: GetTasks :many

SELECT * FROM task ORDER BY id DESC;

-- name: UpdateTask :exec

UPDATE task SET title = $1, description = $2, completed = $3, updated_at = $4 WHERE id = $5;

-- name: DeleteTask :exec

DELETE FROM task WHERE id = $1;

-- name: DeleteAllTasks :exec

DELETE FROM task;



