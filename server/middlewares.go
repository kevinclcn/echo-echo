package server

import (
	"github.com/labstack/echo/middleware"
	"github.com/kevinclcn/log4e"
)

func setupMiddlewares()  {
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:log4e.Output(),
	}))
}
