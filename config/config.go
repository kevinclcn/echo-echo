package config

import (
	"github.com/spf13/viper"
	"fmt"
	"strings"
)

var vip = viper.New()

func init() {
	vip.SetConfigName("application")
	vip.SetConfigType("toml")
	vip.AddConfigPath("./")
	vip.ReadInConfig()
	vip.SetEnvPrefix("echo")
	replacer := strings.NewReplacer(".", "_")
	vip.SetEnvKeyReplacer(replacer)
	vip.BindEnv("server.port")
	vip.SetDefault("server.port", "1213")

}

type server struct {
	Port int
}

func (s server)Address() string  {
	return fmt.Sprintf(":%d", s.Port)
}

type Config struct {
	Server server
}

func LoadConfig() (*Config, error)  {
	var c Config
	err := vip.Unmarshal(&c)
	return &c, err
}
