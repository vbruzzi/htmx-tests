package main

import (
	"context"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
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

func initRouter() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = getTemplates()
	e.Static("/static", "assets")
	handlers.SetupRoutes(e)
	return e.Start(":8080")
}

func connectDb() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")

	if err != nil {
		return err
	}
	defer conn.Close(ctx)
}

func main() {
	err := initRouter()
	if err != nil {
		panic(err)
	}

	err = connectDb()
	if err != nil {
		panic(err)
	}

}
