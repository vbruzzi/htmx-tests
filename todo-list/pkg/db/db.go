package db

import (
	"context"
	"fmt"
	db "vbruzzi/todo-list/db/sqlc"
	"vbruzzi/todo-list/pkg/config"

	"github.com/jackc/pgx/v5"
)

type Db struct {
	Queries *db.Queries
	Close   func()
}

func ConnectDb(config *config.Conf) (*Db, error) {

	ctx := context.Background()
	// todo: read conn from env
	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s",
			config.Db.Host,
			config.Db.Username,
			config.Db.Name,
			config.Db.Pw,
		),
	)

	if err != nil {
		return nil, err
	}

	queries := db.New(conn)

	return &Db{queries, func() { conn.Close(ctx) }}, nil
}
