package routehandler

import (
	"context"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	db "vbruzzi/todo-list/db/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

func parseTemplates() *template.Template {
	templates := template.New("")
	err := filepath.Walk("./views", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templates.ParseFiles(path)
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return templates
}

func getTemplates() *Templates {
	return &Templates{
		templates: parseTemplates(),
	}
}

func setupRoutes(e *echo.Echo, q *db.Queries) {
	e.GET("/", func(c echo.Context) error {
		todos, err := q.ListTodos(context.Background())

		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "index", todos)
	})

	NewListHandlers(e, q)
}

func InitRouter(queries *db.Queries) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = getTemplates()
	e.Static("/static", "assets")
	setupRoutes(e, queries)
	return e.Start(":8080")
}
