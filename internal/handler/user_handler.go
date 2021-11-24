package handler

import (
	"fmt"
	"github.com/devlibx/gox-base/config"
	"go.uber.org/fx"
	"net/http"
)

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
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called AddUserHandler", uh.appConfig.AppName)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}

func (uh *UserHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called GetUserHandler - ", uh.appConfig.AppName)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}
