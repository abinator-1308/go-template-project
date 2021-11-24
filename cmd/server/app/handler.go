package app

import (
	"fmt"
	"github.com/gorilla/mux"
	service "github.com/harishb2k/go-template-project/pkg/core/service"
	"github.com/harishb2k/go-template-project/pkg/server"
	"net/http"
)

func Register(helper *service.Helper) {
	fmt.Println("Got helper in the handler.Register() ", helper.Name)
}

func NewHandler(server server.Server) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}).Methods("GET")

	server.GetRouter()
	return router
}
