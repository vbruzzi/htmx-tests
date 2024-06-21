
-- name: ListUsers :many
SELECT id, username,date_created FROM users;

-- name: SetUserKey :exec
UPDATE users SET login_key = $1 WHERE id = $2;

-- name: GetUserByCredentials :one
SELECT id,login_key FROM users WHERE username = $1 AND password = $2;

-- name: GetUserIdFromLoginKey :one
SELECT id FROM users WHERE login_key = $1;

-- name: GetUserById :one
SELECT id, username, date_created FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, password, date_created) VALUES ($1, $2, NOW()) RETURNING id, username, date_created;
