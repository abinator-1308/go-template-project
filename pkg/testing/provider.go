package testing

import (
	"github.com/devlibx/gox-base"
	"go.uber.org/fx"
)

var TestCommonModule = fx.Options(
	fx.Provide(gox.NewNoOpCrossFunction),
)
