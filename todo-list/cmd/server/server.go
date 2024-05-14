package main

import (
	"vbruzzi/todo-list/pkg/db"
	routehandler "vbruzzi/todo-list/server/routeHandler"
)

func main() {
	queries, close, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	defer close()

	err = routehandler.InitRouter(queries)
	if err != nil {
		panic(err)
	}

}
