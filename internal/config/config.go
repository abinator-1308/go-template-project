package config

import (
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-http/command"
)

type ApplicationConfig struct {
	App          config.App     `json:"app" yaml:"app"`
	Logger       config.Logger  `json:"logger" yaml:"logger"`
	ServerConfig command.Config `json:"server_config" yaml:"server_config"`
}
