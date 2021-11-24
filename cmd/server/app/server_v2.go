package app

import (
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"net/http"
)

type ServerImpl struct {
	fx.In
	server.Server
	AddUserHandler http.HandlerFunc `name:"AddUserHandler"`
	GetUserHandler http.HandlerFunc `name:"GetUserHandler"`
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.GetRouter().ServeHTTP(w, r)
}


func (s *ServerImpl) routes() {
	serviceName := "srv"
	publicRouter := s.GetRouter().Group(serviceName)
	internalRouter := s.GetRouter().Group(serviceName + "/internal")
	_ = internalRouter

	v1Apis := publicRouter.Group("/v1")
	{
		usersApi := v1Apis.Group("users")
		usersApi.POST("", s.handleAddUser())
		usersApi.GET("", s.handleGetUser())
	}
}

func (s *ServerImpl) handleAddUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		s.AddUserHandler.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.GetUserHandler.ServeHTTP(c.Writer, c.Request)
	}
}
