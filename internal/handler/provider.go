package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/harishb2k/go-template-project/internal/common"
	memory "github.com/harishb2k/go-template-project/pkg/database/inmemory"
	"go.uber.org/fx"
)

var TestModule = fx.Options(
	fx.Provide(memory.NewUserRepository),
	fx.Provide(func(repository *memory.UserRepository) common.UserStore { return repository }),
	fx.Provide(gox.NewNoOpCrossFunction),
	fx.Supply(config.App{}),
)

var TestUserHandlerModule = fx.Options(
	UserHandlerModule,
	TestModule,
)
