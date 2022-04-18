package auth

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/sg/routers"
)

func c(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	auth(authRouter)
	// 鉴权中间件
	router.Use(routers.TokenAuthMiddleware())
	userRouter := router.Group("/user")

	user(userRouter)
}

func auth(router *gin.RouterGroup) {
	router.Use(routers.RateLimitMiddleware(1*time.Second, 1, 1))
	router.POST("/token", nil)
}

func user(router *gin.RouterGroup) {
	router.Use(routers.RateLimitMiddleware(1*time.Second, 3, 1))
	router.POST("/who", nil)
}
