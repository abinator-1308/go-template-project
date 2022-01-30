package config

import (
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-http/command"
	messaging "github.com/devlibx/gox-messaging"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
)

type ApplicationConfig struct {
	App             config.App              `json:"app" yaml:"app"`
	Logger          config.Logger           `json:"logger" yaml:"logger"`
	ServerConfig    command.Config          `json:"server_config" yaml:"server_config"`
	MessagingConfig messaging.Configuration `json:"messaging_config" yaml:"messaging_config"`
	MetricConfig    metrics.Config          `json:"metric" yaml:"metric"`
	DynamoConfig    dynamodb.DynamoConfig   `json:"ddb" yaml:"ddb"`
}
