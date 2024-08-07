// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, date_created) VALUES ($1, $2, NOW()) RETURNING id, username, date_created
`

type CreateUserParams struct {
	Username string
	Password string
}

type CreateUserRow struct {
	ID          int32
	Username    string
	DateCreated pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Password)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Username, &i.DateCreated)
	return i, err
}

const getUserByCredentials = `-- name: GetUserByCredentials :one
SELECT id,login_key FROM users WHERE username = $1 AND password = $2
`

type GetUserByCredentialsParams struct {
	Username string
	Password string
}

type GetUserByCredentialsRow struct {
	ID       int32
	LoginKey pgtype.Text
}

func (q *Queries) GetUserByCredentials(ctx context.Context, arg GetUserByCredentialsParams) (GetUserByCredentialsRow, error) {
	row := q.db.QueryRow(ctx, getUserByCredentials, arg.Username, arg.Password)
	var i GetUserByCredentialsRow
	err := row.Scan(&i.ID, &i.LoginKey)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, date_created FROM users WHERE id = $1
`

type GetUserByIdRow struct {
	ID          int32
	Username    string
	DateCreated pgtype.Timestamp
}

func (q *Queries) GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(&i.ID, &i.Username, &i.DateCreated)
	return i, err
}

const getUserIdFromLoginKey = `-- name: GetUserIdFromLoginKey :one
SELECT id FROM users WHERE login_key = $1
`

func (q *Queries) GetUserIdFromLoginKey(ctx context.Context, loginKey pgtype.Text) (int32, error) {
	row := q.db.QueryRow(ctx, getUserIdFromLoginKey, loginKey)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username,date_created FROM users
`

type ListUsersRow struct {
	ID          int32
	Username    string
	DateCreated pgtype.Timestamp
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(&i.ID, &i.Username, &i.DateCreated); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setUserKey = `-- name: SetUserKey :exec
UPDATE users SET login_key = $1 WHERE id = $2
`

type SetUserKeyParams struct {
	LoginKey pgtype.Text
	ID       int32
}

func (q *Queries) SetUserKey(ctx context.Context, arg SetUserKeyParams) error {
	_, err := q.db.Exec(ctx, setUserKey, arg.LoginKey, arg.ID)
	return err
}
