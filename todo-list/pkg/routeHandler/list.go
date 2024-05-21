package routehandler

import (
	"net/http"
	errors "vbruzzi/todo-list/pkg/error"
	todoservice "vbruzzi/todo-list/pkg/todoService"

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
		id,
		todo,
		false,
	}
}

func NewListHandlers(r *Router) {
	g := r.echo.Group("/list")
	service := todoservice.NewTodoService(r.queries)

	g.GET("", func(c echo.Context) error {
		todos, err := service.ListTodos()

		if err != nil {
			return err.Err
		}

		return c.Render(http.StatusOK, "todoList", todos)
	})

	g.POST("", func(c echo.Context) error {
		todo := c.FormValue("value")
		newEntry, err := service.CreateTodo(todo)

		if err.Code == errors.EINVALID {
			res := Form{
				newTodo(0, todo),
				FormErrors{"Value cannot be empty"},
			}

			return c.Render(err.Status, "todoForm", res)
		}

		c.Render(http.StatusOK, "todoForm", nil)
		return c.Render(http.StatusOK, "oobItem", newEntry)
	})

}
