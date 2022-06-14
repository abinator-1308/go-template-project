package command

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/serialization"
	"github.com/harishb2k/go-template-project/internal/common"
	"github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/harishb2k/go-template-project/pkg/database/mysql"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type None string

// NewApplicationEntryPoint is the main entry point function - which will start the server
func NewApplicationEntryPoint(lc fx.Lifecycle, serverImpl ServerImpl) None {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				// First setup routs
				serverImpl.routes()

				// Start server
				if err := serverImpl.Start(); err != nil {
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

// MainWithConfigAsString is used in tests where we have some fixed config
func MainWithConfigAsString(ctx context.Context, configAsString string) *config.ApplicationConfig {
	appConfig := config.ApplicationConfig{}
	err := serialization.ReadYamlFromString(configAsString, &appConfig)
	if err != nil {
		panic(err)
	}
	return Main(ctx, &appConfig)
}

// Main functions starts the app
func Main(ctx context.Context, appConfig *config.ApplicationConfig) *config.ApplicationConfig {

	var dbModule fx.Option
	useDynamoDbForPersistence := true
	if useDynamoDbForPersistence {
		dbModule = fx.Options(
			fx.Provide(func(repository *dynamodb.UserRepository) common.UserStore { return repository }),
			dynamodb.DatabaseModule,
		)
	} else {
		dbModule = fx.Options(
			fx.Provide(func(repository *mysql.UserRepository) common.UserStore { return repository }),
			mysql.DatabaseModule,
		)
	}

	app := fx.New(
		fx.Provide(NewCrossFunctionProvider),
		fx.Invoke(NewApplicationEntryPoint),

		// Setup server
		fx.Provide(server.NewServer),

		// Handlers
		handler.IntegrationModule,
		dbModule,
		// NOTE - FOR MYSQL delete above and add mysql.DatabaseModule

		// Integration module provides some basic capabilities e.g. messaging, metric, and http support
		bootstrap.IntegrationModule,
		fx.Supply(appConfig.App),
		fx.Supply(&appConfig.MessagingConfig),
		fx.Supply(&appConfig.MetricConfig),
		fx.Supply(&appConfig.ServerConfig),
		fx.Supply(&appConfig.DynamoConfig),
		fx.Supply(&appConfig.MySQLConfig),
	)

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

	return appConfig
}

func NewCrossFunctionProvider(metric metrics.Scope) gox.CrossFunction {
	var loggerConfig zap.Config
	loggerConfig = zap.NewDevelopmentConfig()
	logger, _ := loggerConfig.Build()
	return gox.NewCrossFunction(logger, metric)
}
