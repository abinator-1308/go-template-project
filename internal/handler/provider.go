package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/internal/common"
	memory "github.com/harishb2k/go-template-project/pkg/database/inmemory"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
)

// TestModule provides all the basic dependencies for testing handlers
var TestModule = fx.Options(
	fx.Provide(memory.NewUserRepository),
	fx.Provide(func(repository *memory.UserRepository) common.UserStore { return repository }),
	fx.Provide(gox.NewNoOpCrossFunction),
	fx.Supply(config.App{}),
	fx.Provide(func() *gin.Engine {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.Use(server.GinContextToContextMiddleware())
		return r
	}),
)

// TestUserHandlerModule defines all the dependencies needed to test a user handler
var TestUserHandlerModule = fx.Options(
	UserHandlerModule,
	TestModule,
)
