package routes

import (
	"github.com/gin-gonic/gin"
	"vending/app/middleware"
	"vending/app/routes/business"
)

func InitRoute(router *gin.Engine) error {
	router.Use(middleware.LogMiddleware())
	v1 := router.Group("/v1")
	err := initApiRoute(v1)
	return err
}

func initApiRoute(router *gin.RouterGroup) error {
	//业务逻辑路由
	routerAction := router.Group("/api")
	// 登陆
	business.InitBusinessRoute(routerAction)

	return nil
}
