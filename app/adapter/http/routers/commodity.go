package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/handlers/business"
)

func initCommodityRoute(router *gin.RouterGroup, handler *business.CommodityHandler) {
	// 鉴权中间件
	router.Use(handler.TokenAuthMiddleware())
	router = router.Group("/commodity")
	comCmd(router, handler)
	comQuery(router, handler)
}
func comCmd(router *gin.RouterGroup, handler *business.CommodityHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.PUT("/", handler.CreateCommodity)
	router.POST("/", handler.UpdateCommodity)
	router.DELETE("/ids", handler.RemoveCommodity)
}

func comQuery(router *gin.RouterGroup, handler *business.CommodityHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.GET("/:id", handler.GetCommodity)
	router.GET("/up/:id", handler.UpCommodity)
	router.GET("/down/:id", handler.DownCommodity)
	router.POST("/list", handler.ListCommodity)
}
