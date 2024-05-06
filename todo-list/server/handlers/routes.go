package handlers

import (
	"context"
	"net/http"
	db "vbruzzi/todo-list/db/sqlc"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, q *db.Queries) {
	e.GET("/", func(c echo.Context) error {
		todos, err := q.ListTodos(context.Background())

		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "index", todos)
	})

	NewListHandlers(e, q)
}
