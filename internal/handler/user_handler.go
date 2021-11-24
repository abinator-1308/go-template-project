package handler

import (
	"github.com/devlibx/gox-base/config"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"net/http"
	"sync"
)

// UserHandlerModule has all the HTTP hap handlers for user modules. By taking this approach we are able to:
// 1. encapsulate all handlers for user
// 2. User Handler can get everything injected and all handlers can use those dependencies
// 3. Since we return plain HTTP handler, it can be ued by any framework (however you can have Gin specific code here)
var UserHandlerModule = fx.Options(
	fx.Provide(func(appConfig config.App) *UserHandler {
		return &UserHandler{
			appConfig: appConfig,
		}
	}),
)

type UserHandler struct {
	appConfig config.App
}

func (uh *UserHandler) Adduser() http.HandlerFunc {
	initOnce := sync.Once{}
	return func(w http.ResponseWriter, r *http.Request) {

		// One time Add user specific logic goes here
		initOnce.Do(func() {
			// You can do some setup which is very specific to Add user end-point here
		})

		// Get gin context if you want to use
		ginContext := server.GinContextFromHttpRequestVerified(r)
		_ = ginContext

		// Do your logic
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}

func (uh *UserHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get gin context if you want to use
		ginContext := server.GinContextFromHttpRequestVerified(r)
		_ = ginContext

		// Do your logic
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}
