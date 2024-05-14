package routehandler

import (
	"context"
	"fmt"
	"net/http"
	db "vbruzzi/todo-list/db/sqlc"

	"github.com/labstack/echo/v4"
)

type List struct{}

type FormErrors struct {
	Todo string
}

// todo: remove id, make another class
type FormValues struct {
	Id   int
	Todo string
	Done bool
}

type Form struct {
	Values FormValues
	Errors FormErrors
}

func newTodo(id int, todo string) FormValues {
	return FormValues{
		1,
		todo,
		false,
	}
}

func NewListHandlers(e *echo.Echo, q *db.Queries) {
	g := e.Group("/list")

	g.GET("", func(c echo.Context) error {
		todos, err := q.ListTodos(context.Background())

		if err != nil {
			return err
		}

		if err != nil {
			fmt.Println(err)
		}

		return c.Render(http.StatusOK, "todoList", todos)
	})

	g.POST("", func(c echo.Context) error {
		todo := c.FormValue("value")

		if todo == "" {
			res := Form{
				newTodo(0, todo),
				FormErrors{"Value cannot be empty"},
			}

			return c.Render(http.StatusUnprocessableEntity, "todoForm", res)
		}

		newTodo, err := q.CreateTodo(context.Background(), todo)

		if err != nil {
			return err
		}

		c.Render(http.StatusOK, "todoForm", nil)
		return c.Render(http.StatusOK, "oobItem", newTodo)

	})

}
