-- name: Signup :one

INSERT INTO "user" (id, username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: Signin :one

SELECT * FROM "user" WHERE email = $1 and password = $2 LIMIT 1;

-- name: DeleteUserByEmail :one

DELETE FROM "user" WHERE email = $1 RETURNING *;