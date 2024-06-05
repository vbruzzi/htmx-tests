
-- name: ListUsers :many
SELECT id, username, date_created FROM users;

-- name: GetUser :one
SELECT id, username, date_created FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, password, date_created) VALUES ($1, $2, NOW()) RETURNING id, username, date_created;
