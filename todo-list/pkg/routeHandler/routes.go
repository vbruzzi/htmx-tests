package routehandler

import (
	"context"
	"net/http"
	"path/filepath"
	db "vbruzzi/todo-list/db/sqlc"
	templateparser "vbruzzi/todo-list/pkg/templateParser"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	queries *db.Queries
	echo    *echo.Echo
}

func (r *Router) Init() error {
	r.echo.Static("/static", "assets")
	r.echo.GET("/", func(c echo.Context) error {
		todos, err := r.queries.ListTodos(context.Background())

		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "index", todos)
	})

	NewListHandlers(r)

	return r.echo.Start(":8080")
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
