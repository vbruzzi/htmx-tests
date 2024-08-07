-- name: ListTodos :many
SELECT id, todo, created_on FROM todos;

-- name: CreateTodo :one
INSERT INTO todos (todo, created_on) VALUES ($1, NOW()) RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
