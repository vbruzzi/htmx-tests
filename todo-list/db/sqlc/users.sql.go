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

const getUser = `-- name: GetUser :one
SELECT id, username, date_created FROM users WHERE id = $1
`

type GetUserRow struct {
	ID          int32
	Username    string
	DateCreated pgtype.Timestamp
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.Username, &i.DateCreated)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, date_created FROM users
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
