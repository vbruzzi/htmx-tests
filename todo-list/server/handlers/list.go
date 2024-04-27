package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Item struct {
	Value string
}

func listSubroutes(e *echo.Echo) {
	g := e.Group("/list")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "etst")
	})

	g.POST("", func(c echo.Context) error {
		val := c.FormValue("todoItem")

		c.Render(http.StatusOK, "user-input", nil)
		return c.Render(http.StatusOK, "oob-item", Item{val})
	})

}
