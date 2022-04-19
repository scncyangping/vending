package identity

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/sg/handlers/indentity"
)

func InitAuthRoute(router *gin.RouterGroup) {
	handler := indentity.NewAuthHandler()
	authRouter := router.Group("/auth")
	auth(authRouter, handler)
	// 鉴权中间件
	router.Use(handler.TokenAuthMiddleware())
	userRouter := router.Group("/user")

	user(userRouter, handler)
}

func auth(router *gin.RouterGroup, handler *indentity.AuthHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 1, 1))
	router.POST("/login", handler.Login)
}

func user(router *gin.RouterGroup, handler *indentity.AuthHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 3, 1))
	router.POST("/who", handler.Who)
}
