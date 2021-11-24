package app

import (
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServerCommand() *cobra.Command {
	var s server.Server
	injector := fx.New(

		// Main entry point for server
		fx.Invoke(NewApplicationEntryPoint),

		// Register all HTTP API handlers
		handler.Module,
		handler.UserHandlerModule,

		// Basic dependency - underlying server, CrossFunc, configs for application
		fx.Provide(server.NewServer),
		fx.Provide(gox.NewNoOpCrossFunction),
		fx.Supply(config.App{
			AppName:     "example",
			HttpPort:    8098,
			Environment: "test",
		}),

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
