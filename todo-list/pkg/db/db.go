package db

import (
	"context"
	db "vbruzzi/todo-list/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type Db struct {
	Queries *db.Queries
	Close   func()
}

func ConnectDb() (*Db, error) {
	ctx := context.Background()
	// todo: read conn from env
	conn, err := pgx.Connect(ctx, "host=postgres user=postgres dbname=todo_app password=postgres")

	if err != nil {
		return nil, err
	}

	queries := db.New(conn)

	return &Db{queries, func() { conn.Close(ctx) }}, nil
}
