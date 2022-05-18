package routers

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/adapter/http/routers/identity"
	"vending/app/adapter/http/server"
)

func InitRoute(router *gin.Engine, h *server.Handlers) {
	handler := handlers.NewHandler()
	// 加日志中间件
	router.Use(handler.LogMiddleware())
	v1Auth := router.Group("/v1/base")
	identity.InitAuthRoute(v1Auth, h.AuthHandler)
}
