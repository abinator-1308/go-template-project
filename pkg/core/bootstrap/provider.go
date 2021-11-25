package bootstrap

import (
	"go.uber.org/fx"
)

// Module defines all bootstraps which are common for most of the services
// 1. Gox-Http to make HTTP calls to other services
// 2. ...
var Module = fx.Options(
	fx.Invoke(NewMessagingFactory),
	fx.Provide(NewGoxHttpBuilder),
	fx.Provide(NewMetricService),
)
