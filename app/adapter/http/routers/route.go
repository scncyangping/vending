package routers

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/adapter/http/server"
)

func InitRoute(router *gin.Engine, h *server.Handlers) {
	handler := handlers.NewHandler()
	// 加日志中间件
	router.Use(handler.LogMiddleware())
	v1Auth := router.Group("/v1/base")
	initAuthRoute(v1Auth, h.AuthHandler)
	initCommodityRoute(v1Auth, h.CommodityHandler)
	initOrderRoute(v1Auth, h.OrderHandler)
	initInventoryRoute(v1Auth, h.InventoryHandler)
}
