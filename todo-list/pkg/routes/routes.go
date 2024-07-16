package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

const (
	cookieName = "auth"
)

func (r *Router) Init() error {
	realmName := "myrealm"
	clientId := "myclient"
	authority := "localhost:8080"
	base := "/realms/" + realmName + "/protocol/openid-connect"
	redirectUri := "http://localhost:3000/login"

	r.echo.GET("/login", func(c echo.Context) error {
		if code := c.Request().URL.Query().Get("code"); code != "" {
			redir := url.URL{}
			redir.Scheme = "http"
			redir.Host = "identity:8080"
			redir.Path = base + "/token"

			data := url.Values{}
			data.Add("client_id", "myclient")
			data.Add("grant_type", "authorization_code")
			data.Add("code", code)
			data.Add("redirect_uri", redirectUri)

			res, _ := http.PostForm(redir.String(), data)

			defer res.Body.Close()

			v := map[string]string{}
			err := json.NewDecoder(res.Body).Decode(&v)

			if err != nil {
				fmt.Print("err3")
				fmt.Println(err)
			}

			fmt.Printf("%+v\n", v["access_token"])

			c.SetCookie(&http.Cookie{
				Name:  cookieName,
				Value: v["access_token"],
			})

			c.Redirect(http.StatusSeeOther, "/")

			return nil
		}

		if _, err := c.Cookie(cookieName); err != nil {
			redir := url.URL{}
			redir.Scheme = "http"
			redir.Host = authority
			redir.Path = base + "/auth"

			query := redir.Query()
			query.Add("response_type", "code")
			query.Add("scope", "openid email")
			query.Add("state", "random")
			query.Add("client_id", clientId)
			query.Add("redirect_uri", redirectUri)

			redir.RawQuery = query.Encode()

			c.Redirect(http.StatusSeeOther, redir.String())
			return nil
		}

		c.Redirect(http.StatusSeeOther, "/")
		return nil
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
