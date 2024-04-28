package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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

func listSubroutes(e *echo.Echo) {
	g := e.Group("/list")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "etst")
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

		c.Render(http.StatusOK, "todoForm", nil)
		return c.Render(http.StatusOK, "oobItem", newTodo(1, todo))

	})

}
