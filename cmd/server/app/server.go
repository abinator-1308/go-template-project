package app

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/core/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"net/http"
)

type ServerImpl struct {
	fx.In
	server.Server

	UserHandler   *handler.UserHandler
	MetricHandler *bootstrap.MetricHandler
	Cf            gox.CrossFunction
	AppConfig     config.App

	RandomApiHandler          http.HandlerFunc `name:"RandomApiHandler"`
	JsonPlaceholderApiHandler http.HandlerFunc `name:"JsonPlaceholderApiHandler"`
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.GetRouter().ServeHTTP(w, r)
}

func (s *ServerImpl) routes() {
	serviceName := s.AppConfig.AppName
	publicRouter := s.GetRouter().Group(serviceName)
	internalRouter := s.GetRouter().Group(serviceName + "/internal")

	// We must add the MW with all the routers
	s.GetRouter().Use(server.GinContextToContextMiddleware())
	publicRouter.Use(server.GinContextToContextMiddleware())
	internalRouter.Use(server.GinContextToContextMiddleware())

	/*s.GetRouter().Use(ginhttp.Middleware(opentracing.GlobalTracer()))
	publicRouter.Use(ginhttp.Middleware(opentracing.GlobalTracer()))
	internalRouter.Use(ginhttp.Middleware(opentracing.GlobalTracer()))*/

	s.GetRouter().Use(gintrace.Middleware(serviceName))
	publicRouter.Use(gintrace.Middleware(serviceName))
	internalRouter.Use(gintrace.Middleware(serviceName))

	// register metric
	{
		publicRouter.GET("/metrics", s.handleMetricRequest())
	}

	// All v1 APIS
	v1Apis := publicRouter.Group("/v1")
	{
		// User APIs
		usersApi := v1Apis.Group("users")
		usersApi.POST("", s.handleAddUser())
		usersApi.GET("", s.handleGetUser())
	}
	{
		// Random APIs
		randomApi := v1Apis.Group("random")
		randomApi.GET("", s.handleRandomApi())
		randomApi.GET("/JsonPlaceholderApiHandler", s.handleJsonPlaceholderApiHandler())
	}
}

func (s *ServerImpl) handleMetricRequest() func(c *gin.Context) {
	return func(c *gin.Context) {
		s.MetricHandler.HTTPHandler().ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleAddUser() func(c *gin.Context) {
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

func (s *ServerImpl) handleRandomApi() func(c *gin.Context) {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.RandomApiHandler).ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleJsonPlaceholderApiHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.JsonPlaceholderApiHandler).ServeHTTP(c.Writer, c.Request)
	}
}
