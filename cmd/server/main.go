package main

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/serialization"
	app2 "github.com/harishb2k/go-template-project/cmd/server/app"
	"github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/clients"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/harishb2k/go-template-project/pkg/server"
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
		app2.Module,

		// Setup server
		fx.Provide(server.NewServer),

		// Handlers
		handler.IntegrationModule,
		dynamodb.DatabaseModule,

		// Integration module provides some basic capabilities e.g. messaging, metric, and http support
		bootstrap.IntegrationModule,
		fx.Supply(appConfig.App),
		fx.Supply(&appConfig.MessagingConfig),
		fx.Supply(&appConfig.MetricConfig),
		fx.Supply(&appConfig.ServerConfig),
		fx.Supply(&appConfig.DynamoConfig),

		// Integration module to get all clients
		clients.IntegrationModule,
		fx.Populate(&jsonPlaceholderClient),
	)

	ctx := context.Background()
	err = app.Start(ctx)
	if err != nil {
		panic(err)
	}
	<-ctx.Done()
}

func NewCrossFunctionProvider(metric metrics.Scope) gox.CrossFunction {
	var loggerConfig zap.Config
	loggerConfig = zap.NewDevelopmentConfig()
	logger, _ := loggerConfig.Build()
	return gox.NewCrossFunction(logger, metric)
}
