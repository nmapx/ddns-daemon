package config

import (
	elog "github.com/labstack/gommon/log"
	"github.com/nmapx/ddns-daemon/file"
	"gopkg.in/yaml.v2"
	"os"
)

//Config struct
type Config struct {
	Hosts map[string]HostConfig `yaml:"hosts"`
}

// HostConfig struct
type HostConfig struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// Load config file
func (c *Config) Load(filepath string) {
	if !file.Exists(filepath) {
		elog.Errorf("Config file does not exist")
		os.Exit(1)
	}

	err := yaml.Unmarshal(file.Read(filepath), &c)
	if err != nil {
		elog.Errorf("Unable to read config: %v", err)
		os.Exit(1)
	}
}
