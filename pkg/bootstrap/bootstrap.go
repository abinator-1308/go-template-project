package bootstrap

import (
	"context"
	"go.uber.org/fx"
)

type None string

func NewBootstrapStartup(lc fx.Lifecycle, factory MessagingFactory) None {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return ""
}
