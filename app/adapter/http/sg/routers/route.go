package routers

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/sg/handlers"
	"vending/app/adapter/http/sg/routers/identity"
)

func InitRoute(router *gin.Engine) {
	handler := handlers.NewHandler()
	// 加日志中间件
	router.Use(handler.LogMiddleware())
	v1 := router.Group("/v1")
	identity.InitAuthRoute(v1)
}
