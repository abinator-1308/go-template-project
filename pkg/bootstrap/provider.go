package bootstrap

import "go.uber.org/fx"

var IntegrationModule = fx.Options(
	fx.Invoke(NewMessagingFactory),
	fx.Provide(NewMessagingFactory),
	fx.Provide(NewGoxHttpBuilder),
	fx.Provide(NewMetricService),
)
