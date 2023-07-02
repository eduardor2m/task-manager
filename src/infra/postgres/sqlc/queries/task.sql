-- name: CreateTask :exec

INSERT INTO task (
    user_id,
    id,
    title,
    description,
    category,
    status,
    date,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) RETURNING *;

-- name: GetTask :one

SELECT * FROM task WHERE id = $1;

-- name: GetTasks :many

SELECT * FROM task WHERE user_id = $1 ORDER BY id;

-- name: UpdateTask :one

UPDATE task SET title = $1, description = $2, category = $3, status = $4, date = $5, updated_at = $6 WHERE id = $7 RETURNING *;

-- name: UpdateTaskStatus :one

UPDATE task SET status = $1, updated_at = $2 WHERE id = $3 RETURNING *;

-- name: DeleteTask :exec

DELETE FROM task WHERE id = $1;

-- name: DeleteAllTasks :exec

DELETE FROM task WHERE user_id = $1;



