-- name: CreateTask :exec

INSERT INTO task (title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?);

-- name: GetTask :one

SELECT * FROM task WHERE id = ?;

-- name: UpdateTask :exec

UPDATE task SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?;

-- name: DeleteTask :exec

DELETE FROM task WHERE id = ?;

