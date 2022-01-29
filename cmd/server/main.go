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

		fx.Provide(bootstrap.NewMetricService),
		fx.Supply(appConfig.App),
		fx.Supply(&appConfig.MetricConfig), // For messaging (if you don't use messaging pass a object with metric config enabled = false)

		fx.Provide(bootstrap.NewGoxHttpBuilder),
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
