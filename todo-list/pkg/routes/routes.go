package routes

import (
	"context"
	"net/http"
	"path/filepath"
	db "vbruzzi/todo-list/db/sqlc"
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
	Todos []db.Todo
}

type AuthHandler interface {
	Authenticate(c echo.Context) error
	IsAuthenticated(c echo.Context) bool
}

func (r *Router) Init(authHandler AuthHandler) error {
	r.echo.Static("/static", "assets")

	r.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() == "/login" || authHandler.IsAuthenticated(c) {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login")
		}
	})

	r.echo.GET("/login", authHandler.Authenticate)

	r.echo.GET("/", func(c echo.Context) error {
		data := homeData{}

		_, err := r.queries.ListTodos(context.Background())

		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "index", data)
	})

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
