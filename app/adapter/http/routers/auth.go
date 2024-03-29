package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/handlers/business"
)

func initAuthRoute(router *gin.RouterGroup, handler *business.AuthHandler) {
	auth(router, handler)
	// 鉴权中间件
	router.Use(handler.TokenAuthMiddleware())
	userRouter := router.Group("/user")

	user(userRouter, handler)
}

func auth(router *gin.RouterGroup, handler *business.AuthHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 1, 1))
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)

}

func user(router *gin.RouterGroup, handler *business.AuthHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 3, 1))
	router.POST("/who", nil)
}
