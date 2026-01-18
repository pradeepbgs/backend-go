-- name: GetUsers :many
SELECT id,name,email FROM users;

-- name: GetUserById :one
SELECT id,name,email FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name,email)
VALUES ($1,$2) RETURNING *;

-- name: GetUserByEmail :one
SELECT id, name, email FROM users WHERE email = $1;
