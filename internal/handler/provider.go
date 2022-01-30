package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/internal/common"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	memory "github.com/harishb2k/go-template-project/pkg/database/inmemory"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
)

// UserHandlerModule has all the HTTP hap handlers for user modules. By taking this approach we are able to:
// 1. encapsulate all handlers for user
// 2. User Handler can get everything injected and all handlers can use those dependencies
// 3. Since we return plain HTTP handler, it can be ued by any framework (however you can have Gin specific code here)
var UserHandlerModule = fx.Options(
	fx.Provide(func(cf gox.CrossFunction, appConfig config.App, userDao common.UserStore) *UserHandler {
		return &UserHandler{
			appConfig: appConfig,
			cf:        cf,
			userDao:   userDao,
		}
	}),
	fx.Provide(func(repository *dynamodb.UserRepository) common.UserStore { return repository }),
)

// IntegrationModule is full wired module to be used in application
var IntegrationModule = fx.Options(
	UserHandlerModule,
)

// -------------------------------------- Testing Modules --------------------------------------------------------------

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
