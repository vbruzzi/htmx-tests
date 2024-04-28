package handlers

import "github.com/labstack/echo/v4"

func todoSubroutes(e *echo.Echo) {
	g := e.Group("/todo")

	g.GET("/:id", func(c echo.Context) error {

	})
}
