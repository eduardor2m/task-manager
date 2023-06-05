-- name: Signup :one

INSERT INTO "user" (id, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: Signin :one

SELECT * FROM "user" WHERE email = $1 LIMIT 1;