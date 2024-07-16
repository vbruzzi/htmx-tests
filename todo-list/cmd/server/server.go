package main

import (
	"log"
	"vbruzzi/todo-list/pkg/auth"
	"vbruzzi/todo-list/pkg/config"
	"vbruzzi/todo-list/pkg/db"
	"vbruzzi/todo-list/pkg/routes"
)

func main() {
	conf, err := config.NewConfig()

	if err != nil {
		log.Fatalf("failed to read config: %+v", err)
	}

	res, err := db.ConnectDb(conf)
	if err != nil {
		panic(err)
	}

	defer res.Close()

	router, err := routes.NewRouter(res.Queries)
	if err != nil {
		panic(err)
	}

	authHandler := auth.New(conf.Oidc)

	router.Init(&authHandler)
}
