package app

import (
	"context"
	"go.uber.org/fx"
)

func NewApplicationEntryPoint(
	lc fx.Lifecycle,
	serverImpl ServerImpl,
) string {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				serverImpl.routes()
				serverImpl.Start()
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
