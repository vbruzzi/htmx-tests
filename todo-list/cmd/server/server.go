package main

import (
	"vbruzzi/todo-list/pkg/db"
	routehandler "vbruzzi/todo-list/pkg/routeHandler"
)

func main() {
	res, err := db.ConnectDb()
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
