package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	Value string
}

type Form struct {
	Values FormData
	Errors FormData
}

func listSubroutes(e *echo.Echo) {
	g := e.Group("/list")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "etst")
	})

	g.POST("", func(c echo.Context) error {
		value := c.FormValue("value")

		if value == "" {
			res := Form{
				FormData{value},
				FormData{"Value cannot be empty"},
			}

			return c.Render(http.StatusUnprocessableEntity, "user-input", res)
		}

		c.Render(http.StatusOK, "user-input", nil)
		return c.Render(http.StatusOK, "oob-item", FormData{value})
	})

}
