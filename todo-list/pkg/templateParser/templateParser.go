package templateparser

import (
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseTemplates() {

}

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

func GetTemplates() *Templates {
	return &Templates{
		templates: parseTemplates(),
	}
}
