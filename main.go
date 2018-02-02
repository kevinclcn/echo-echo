package main

import (
	"github.com/kevinclcn/echo-echo/config"
	"github.com/labstack/gommon/log"
	"github.com/kevinclcn/echo-echo/server"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	server.Start(conf)
}