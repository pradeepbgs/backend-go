-- name: GetUsers :many
SELECT id,name,email FROM users;

-- name: GetUserById :one
SELECT id,name,email FROM users WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (name,email)
VALUES ($1,$2);