package main

import (
	"github.com/kevinclcn/echo-echo/config"
	"github.com/kevinclcn/echo-echo/server"
	"github.com/kevinclcn/log4e"
)

func main() {

	config.SetupLog()
	conf, err := config.LoadConfig()
	if err != nil {
		log4e.Fatal(err)
	}

	log4e.Info("start server")
	server.Start(conf)
}