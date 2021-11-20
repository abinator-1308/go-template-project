package app

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base"
	c "github.com/devlibx/gox-base/config"
	goxServer "github.com/devlibx/gox-base/server"
	common2 "github.com/harishb2k/go-template-project/pkg/common"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
	"go.uber.org/fx"
	"net/http"
)

// NewEntryPoint
// TODO: Comment to help understand code - Delete me
// This is registered as "fx.Provide(NewEntryPoint)", now when this provider is called, then it gets "fx.Lifecycle" in
// input. You can add your hook which is called when app is started
func NewEntryPoint(lc fx.Lifecycle, helper *common.Helper, someConstantIntValueFromConfig int, config *common2.Config,
	handler http.Handler,
) string {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting HTTP server. ", helper.Name, someConstantIntValueFromConfig, config.Env)

			serverInstance, _ := goxServer.NewServer(gox.NewNoOpCrossFunction())
			serverInstance.Start(handler, &c.App{
				AppName:     "dummy",
				HttpPort:    9090,
				Environment: "prod",
			})

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server.", helper.Name)
			return nil
		},
	})
	return ""
}
