package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/handlers/business"
)

func initOrderRoute(router *gin.RouterGroup, handler *business.OrderHandler) {
	// 鉴权中间件
	router.Use(handler.TokenAuthMiddleware())
	router = router.Group("/order")
	orderCmd(router, handler)
	orderQuery(router, handler)
}
func orderCmd(router *gin.RouterGroup, handler *business.OrderHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.PUT("/", handler.CreateOrder)
	router.GET("/call/:id", handler.OrderCallBack)
	router.GET("/cancel/:id", handler.Cancel)

}

func orderQuery(router *gin.RouterGroup, handler *business.OrderHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.GET("/:id", handler.GetOrderById)
	router.GET("/temp/:id", handler.GetTempOrderById)
	router.POST("/list", handler.Query)
}
