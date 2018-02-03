package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func setupUrlMappings()  {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
