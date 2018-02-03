package server

import "github.com/labstack/echo/middleware"

func setupMiddlewares()  {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
}
