package command

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"net/http"
)

type ServerImpl struct {
	fx.In
	server.Server

	// Common properties
	Cf        gox.CrossFunction
	AppConfig config.App

	// Handlers
	UserHandler *handler.UserHandler
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.GetRouter().ServeHTTP(w, r)
}

func (s *ServerImpl) routes() {
	serviceName := s.AppConfig.AppName
	publicRouter := s.GetRouter().Group(serviceName)

	s.GetRouter().Use(server.GinContextToContextMiddleware())
	publicRouter.Use(server.GinContextToContextMiddleware())
	s.GetRouter().Use(gintrace.Middleware(serviceName))
	publicRouter.Use(gintrace.Middleware(serviceName))

	// All v1 APIS
	v1Apis := publicRouter.Group("/v1")
	{
		// User APIs
		usersApi := v1Apis.Group("users")
		usersApi.POST("", s.handleAddUser())
		usersApi.GET("/:id/:property", s.handleGetUser())
	}

}

func (s *ServerImpl) handleAddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		hf := server.EnsureGinContextWrapper(s.UserHandler.Adduser())
		hf = server.MetricWrapper(hf, s.Cf, "add_user")
		hf.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.UserHandler.GetUser()).ServeHTTP(c.Writer, c.Request)
	}
}
