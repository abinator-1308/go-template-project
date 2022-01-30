package app

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/serialization"
	"github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type None string

var Module = fx.Options(
	fx.Invoke(NewApplicationEntryPoint),
)

func NewApplicationEntryPoint(
	lc fx.Lifecycle,
	serverImpl ServerImpl,
) None {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				serverImpl.routes()
				err := serverImpl.Start()
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			<-serverImpl.Stop()
			return nil
		},
	})
	return ""
}

func Main(ctx context.Context, configLocation string) {
	appConfig := config.ApplicationConfig{}
	err := serialization.ReadYaml(configLocation, &appConfig)
	if err != nil {
		panic(err)
	}

	app := fx.New(
		fx.Provide(NewCrossFunctionProvider),
		Module,

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
	)

	err = app.Start(ctx)
	if err != nil {
		panic(err)
	}
}

func NewCrossFunctionProvider(metric metrics.Scope) gox.CrossFunction {
	var loggerConfig zap.Config
	loggerConfig = zap.NewDevelopmentConfig()
	logger, _ := loggerConfig.Build()
	return gox.NewCrossFunction(logger, metric)
}
