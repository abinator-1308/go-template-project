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
