package handler

import (
	"fmt"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"net/http"
)

var RandomHandlerModule = fx.Options(
	fx.Provide(fx.Annotated{Name: "RandomApiHandler", Target: func(server server.Server) http.HandlerFunc {
		return randomApiHandler()
	}}),
)

func randomApiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called random endpoint")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}
