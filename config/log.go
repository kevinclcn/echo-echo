package config

import (
	"github.com/kevinclcn/log4e"
	"github.com/sirupsen/logrus"
)

func SetupLog()  {
	log4e.SetFormatter(&logrus.JSONFormatter{})
}
