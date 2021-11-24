package handler

import (
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(fx.Annotated{Name: "AddUserHandler", Target: Adduser}),
	fx.Provide(fx.Annotated{Name: "GetUserHandler", Target: GetUser}),
)

func Adduser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called AddUserHandler")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called GetUserHandler")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}
}
