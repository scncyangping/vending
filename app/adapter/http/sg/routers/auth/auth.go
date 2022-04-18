package auth

import (
	"github.com/gin-gonic/gin"
	"time"
)

func InitAuthRoute(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	auth(authRouter)
	// 鉴权中间件
	router.Use(TokenAuthMiddleware())
	userRouter := router.Group("/user")

	user(userRouter)
}

func auth(router *gin.RouterGroup) {
	auth := auth3.AuthController{}
	router.Use(RateLimitMiddleware(1*time.Second, 1, 1))
	router.POST("/token", auth.Login)
}

func user(router *gin.RouterGroup) {
	c := auth3.UserController{}
	router.Use(RateLimitMiddleware(1*time.Second, 3, 1))
	router.POST("/who", c.WhoIam)
}
