package app

import (
	"fmt"
	"github.com/gorilla/mux"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
	"net/http"
)

func Register(helper *common.Helper) {
	fmt.Println("Got helper in the handler.Register() ", helper.Name)
}

func NewHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}).Methods("GET")
	return router
}
