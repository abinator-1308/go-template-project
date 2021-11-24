package server

import "github.com/gin-gonic/gin"

type RouteRegistry interface {
	SetupRoute(router *gin.Engine) error
}
