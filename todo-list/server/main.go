package main

import (
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
	"vbruzzi/todo-list/server/handlers"

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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = getTemplates()
	handlers.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":888"))
}
