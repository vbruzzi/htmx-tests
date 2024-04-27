package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	Todo string
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
		todo := c.FormValue("value")

		if todo == "" {
			res := Form{
				FormData{todo},
				FormData{"Value cannot be empty"},
			}

			return c.Render(http.StatusUnprocessableEntity, "todoForm", res)
		}

		c.Render(http.StatusOK, "todoForm", nil)
		return c.Render(http.StatusOK, "oobItem", FormData{todo})
	})

}
