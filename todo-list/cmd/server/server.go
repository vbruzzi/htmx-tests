package main

import (
	"log"
	"vbruzzi/todo-list/pkg/config"
	"vbruzzi/todo-list/pkg/db"
	routehandler "vbruzzi/todo-list/pkg/routeHandler"
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

	router, err := routehandler.NewRouter(res.Queries)
	if err != nil {
		panic(err)
	}

	router.Init()
}
