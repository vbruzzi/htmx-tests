package main

import (
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getRenderer() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.go.tpl")),
	}
}

type Contact struct {
	Name string
	Id   int
}

type ContactBatch struct {
	Data []Contact
	Next int
}

func newContact(name string, id int) Contact {
	return Contact{
		name,
		id,
	}
}

func getContactsList(start int, take int) ContactBatch {
	contacts := []Contact{}

	for i := start; i < start+take; i++ {
		contacts = append(contacts, newContact(RandStringRunes(10), i))
	}

	return ContactBatch{
		contacts,
		start + take,
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = getRenderer()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", getContactsList(0, 10))
	})

	e.GET("/contacts", func(c echo.Context) error {
		take, err := strconv.Atoi(c.QueryParam("take"))

		if err != nil {
			return err
		}

		start, err := strconv.Atoi(c.QueryParam("start"))

		if err != nil {
			return err
		}
		return c.Render(http.StatusOK, "contacts", getContactsList(start, take))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
