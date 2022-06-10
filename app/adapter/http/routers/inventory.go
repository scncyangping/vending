package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"vending/app/adapter/http/handlers/business"
)

func initInventoryRoute(router *gin.RouterGroup, handler *business.InventoryHandler) {
	// 鉴权中间件
	router.Use(handler.TokenAuthMiddleware())
	router = router.Group("/inventory")
	iitCmd(router, handler)
	iitQuery(router, handler)
}
func iitCmd(router *gin.RouterGroup, handler *business.InventoryHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.PUT("/", handler.SaveCategory)
	router.PUT("/stock", handler.InStockOne)
	router.POST("/", handler.UpdateCategory)
	router.DELETE("/ids", handler.RemoveCategoryByIds)

}

func iitQuery(router *gin.RouterGroup, handler *business.InventoryHandler) {
	router.Use(handler.RateLimitMiddleware(1*time.Second, 10, 10))
	router.POST("/list", handler.QueryCategory)
	router.POST("/stock/list", handler.QueryStock)
}
