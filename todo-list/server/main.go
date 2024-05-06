package main

import (
	"context"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
	db "vbruzzi/todo-list/db/sqlc"
	"vbruzzi/todo-list/server/handlers"

	"github.com/jackc/pgx/v5"
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

func initRouter(queries *db.Queries) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = getTemplates()
	e.Static("/static", "assets")
	handlers.SetupRoutes(e, queries)
	return e.Start(":8080")
}

func connectDb() (*db.Queries, func(), error) {
	ctx := context.Background()
	// todo: read conn from env
	conn, err := pgx.Connect(ctx, "host=postgres user=postgres dbname=todo_app password=postgres")

	if err != nil {
		return nil, nil, err
	}

	queries := db.New(conn)

	return queries, func() { conn.Close(ctx) }, nil
}

func main() {
	queries, close, err := connectDb()
	if err != nil {
		panic(err)
	}

	defer close()

	err = initRouter(queries)
	if err != nil {
		panic(err)
	}

}
