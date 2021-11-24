package config

import (
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-http/command"
	messaging "github.com/devlibx/gox-messaging"
)

type ApplicationConfig struct {
	App             config.App              `json:"app" yaml:"app"`
	Logger          config.Logger           `json:"logger" yaml:"logger"`
	ServerConfig    command.Config          `json:"server_config" yaml:"server_config"`
	MessagingConfig messaging.Configuration `json:"messaging_config" yaml:"messaging_config"`
}
