package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"os"
)

func TestConfigLoader(t *testing.T)  {
	var tomlExample = []byte(`
		address = ":1212"

     `)

	vip.ReadConfig(bytes.NewBuffer(tomlExample))

	c, err := LoadConfig()
	assert.Nil(t, err, "failed to load config file")
	assert.Equal(t, ":1212", c.Address, "Address is not equal")

	vip.ReadConfig(bytes.NewBuffer([]byte(``)))
}

func TestEnv(t *testing.T)  {
	os.Setenv("ECHO_ADDRESS", ":1111")
	c, err := LoadConfig()
	assert.Nil(t, err, "failed to load config file")
	assert.Equal(t, ":1111", c.Address, "Address is not equal")

	os.Clearenv()
}

func TestDefault(t *testing.T) {
	c, err := LoadConfig()
	assert.Nil(t, err, "failed to load config file")
	assert.Equal(t, ":1213", c.Address, "Address is not equal")
}