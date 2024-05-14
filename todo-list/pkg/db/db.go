package db

import (
	"context"
	db "vbruzzi/todo-list/db/sqlc"

	"github.com/jackc/pgx/v5"
)

func ConnectDb() (*db.Queries, func(), error) {
	ctx := context.Background()
	// todo: read conn from env
	conn, err := pgx.Connect(ctx, "host=postgres user=postgres dbname=todo_app password=postgres")

	if err != nil {
		return nil, nil, err
	}

	queries := db.New(conn)

	return queries, func() { conn.Close(ctx) }, nil
}
