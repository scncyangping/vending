package business

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/handler/business"
	"vending/app/middleware"
)

/*
 * date : 2019/4/30
 * author : yangping
 * desc : 所有业务模块路由汇总，每个模块具体方法再分文件写
 */
func InitBusinessRoute(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	auth(authRouter)
	// 鉴权中间件
	router.Use(middleware.TokenAuthMiddleware())
	userRouter := router.Group("/user")

	user(userRouter)
}

func auth(router *gin.RouterGroup) {
	auth := business.AuthController{}
	router.Use(middleware.RateLimitMiddleware(1*time.Second, 1, 1))
	router.POST("/token", auth.Login)
}

func user(router *gin.RouterGroup) {
	c := business.UserController{}
	router.Use(middleware.RateLimitMiddleware(1*time.Second, 3, 1))
	router.POST("/who", c.WhoIam)
}
