package app

import (
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/serialization"
	config2 "github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/clients"
	"github.com/harishb2k/go-template-project/pkg/core/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServerCommand() *cobra.Command {

	appConfig := config2.ApplicationConfig{}
	err := serialization.ReadYaml("./config/app.yaml", &appConfig)
	if err != nil {
		panic(err)
	}

	var s server.Server
	injector := fx.New(

		// Main entry point for server
		fx.Invoke(NewApplicationEntryPoint),

		// Bootstrap dependencies - e.g. Gox HTTP, messaging, caching, ...
		bootstrap.Module,

		// Client module - these are the HTTP clients which this service can call
		clients.Module,

		// Basic dependency - underlying server, CrossFunc, configs for application
		fx.Provide(server.NewServer),
		fx.Provide(gox.NewNoOpCrossFunction),
		fx.Supply(config.App{
			AppName:     "example",
			HttpPort:    8098,
			Environment: "test",
		}),

		fx.Supply(&appConfig.ServerConfig), // Gox-Http config which is needed by bootstrap module

		fx.Supply(&appConfig.MessagingConfig),

		// Register all HTTP API handlers
		handler.RandomHandlerModule,
		handler.UserHandlerModule,

		// Instance of underlying server
		fx.Populate(&s),
	)

	return &cobra.Command{
		Use:   "gox",
		Short: "Small help code",
		Long:  `Long help code`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Your service running code here...")
			injector.Run()
		},
	}
}
