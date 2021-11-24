package handler

import (
	"fmt"
	"github.com/devlibx/gox-base/serialization"
	"github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"net/http"
)

// RandomHandlerModule provides example of a random API you can add
var RandomHandlerModule = fx.Options(
	fx.Provide(fx.Annotated{Name: "RandomApiHandler", Target: func(server server.Server) http.HandlerFunc {
		return randomApiHandler()
	}}),
	fx.Provide(fx.Annotated{Name: "JsonPlaceholderApiHandler", Target: func(client jsonplaceholder.Client) http.HandlerFunc {
		return jsonPlaceHolderApiHandler(client)
	}}),
)

func randomApiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called random endpoint")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}

func jsonPlaceHolderApiHandler(client jsonplaceholder.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post, err := client.FetchPost(r.Context(), "1")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(serialization.ToBytesSuppressError(post))
	}
}
