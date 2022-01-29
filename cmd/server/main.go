package main

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/serialization"
	"github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/pkg/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/clients"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	appConfig := config.ApplicationConfig{}
	err := serialization.ReadYaml("./config/app.yaml", &appConfig)
	if err != nil {
		panic(err)
	}

	var jsonPlaceholderClient clients.JsonPlaceholderClient
	app := fx.New(
		fx.Provide(NewCrossFunctionProvider),

		clients.IntegrationJsonPlaceholderModule,
		fx.Populate(&jsonPlaceholderClient),

		// Integration module provides some basic capabilities e.g. messaging, metric, and http support
		bootstrap.IntegrationModule,
		fx.Supply(appConfig.App),
		fx.Supply(&appConfig.MessagingConfig),
		fx.Supply(&appConfig.MetricConfig),
		fx.Supply(&appConfig.ServerConfig),

	)
	err = app.Start(context.Background())
	if err != nil {
		panic(err)
	}

	post, err := jsonPlaceholderClient.FetchPost(context.Background(), "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(post)
}

func NewCrossFunctionProvider(metric metrics.Scope) gox.CrossFunction {
	var loggerConfig zap.Config
	loggerConfig = zap.NewDevelopmentConfig()
	logger, _ := loggerConfig.Build()
	return gox.NewCrossFunction(logger, metric)
}
