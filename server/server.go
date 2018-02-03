package server

import (
	"github.com/labstack/echo"
	"github.com/kevinclcn/echo-echo/config"
)

var e = echo.New()

func Start(config *config.Config)  {
	setupMiddlewares()
	setupUrlMappings()
	setupLogin()

	e.Logger.Fatal(e.Start(config.Server.Address()))
}