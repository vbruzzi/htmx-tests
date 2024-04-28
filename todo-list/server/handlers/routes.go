package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", []FormData{{"Test"}})
	})

	listSubroutes(e)
}
