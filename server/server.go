package server

import (
	"github.com/labstack/echo"
	"github.com/kevinclcn/echo-echo/config"
	"github.com/kevinclcn/log4e"
)

var e = echo.New()

func Start(config *config.Config)  {
	e.Logger = log4e.Logger()
	setupMiddlewares()
	setupUrlMappings()
	setupLogin()

	e.Logger.Fatal(e.Start(config.Server.Address()))
}