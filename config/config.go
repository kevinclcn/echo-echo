package config

import (
	"github.com/spf13/viper"
)

var vip = viper.New()

func init() {
	vip.SetConfigName("application")
	vip.SetConfigType("toml")
	vip.AddConfigPath("./")
	vip.ReadInConfig()
	vip.SetEnvPrefix("echo")
	vip.BindEnv("address")
	vip.SetDefault("address", ":1213")

}

type Config struct {
	Address string
}

func LoadConfig() (*Config, error)  {
	var c Config
	err := vip.Unmarshal(&c)
	return &c, err
}
