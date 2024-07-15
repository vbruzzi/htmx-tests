package routes

import (
	"context"
	"net/http"
	"path/filepath"
	db "vbruzzi/todo-list/db/sqlc"
	"vbruzzi/todo-list/pkg/handlers/login"
	"vbruzzi/todo-list/pkg/handlers/todos"
	templateparser "vbruzzi/todo-list/pkg/templateParser"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	queries *db.Queries
	echo    *echo.Echo
}

type homeData struct {
	Todos    []db.Todo
	LoggedIn bool
}

func (r *Router) Init() error {
	r.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := c.Cookie("auth")

			if err != nil {
				// handle redirect
			}

			if err = next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	})
	r.echo.Static("/static", "assets")
	r.echo.GET("/", func(c echo.Context) error {
		data := homeData{LoggedIn: false}

		_, err := r.queries.ListTodos(context.Background())

		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "index", data)
	})

	loginGroup := r.echo.Group("/login")
	login.NewLoginHandler(loginGroup, r.queries)

	todoGroup := r.echo.Group("/todos")
	todos.NewTodoHandler(todoGroup, r.queries)

	return r.echo.Start(":3000")
}

func NewRouter(queries *db.Queries) (*Router, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templateparser.ParseTemplates(filepath.Walk)
	return &Router{
		echo:    e,
		queries: queries,
	}, nil
}
