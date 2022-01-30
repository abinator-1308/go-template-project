package bootstrap

import (
	"context"
	goxMessaging "github.com/devlibx/gox-messaging"
	"go.uber.org/fx"
)

type None string

func NewBootstrapStartup(lc fx.Lifecycle, messagingFactory MessagingFactory, configuration *goxMessaging.Configuration) None {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if configuration.Enabled {
				return messagingFactory.Start(*configuration)
			} else {
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return messagingFactory.Stop()
		},
	})
	return ""
}
