package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/pkg/server"
)

func setupGinForTesting() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(server.GinContextToContextMiddleware())
	return r
}
