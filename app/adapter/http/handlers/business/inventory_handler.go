package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/service"
)

type InventoryHandler struct {
	*handlers.Handler
	inventorySrv service.InventorySrv
}

func NewInventoryHandler(handler *handlers.Handler, inventorySrv service.InventorySrv) *InventoryHandler {
	return &InventoryHandler{Handler: handler, inventorySrv: inventorySrv}
}

func (i *InventoryHandler) SaveCategory(ctx *gin.Context) {
	var (
		requestBody cmd.CategorySaveCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if rp, err := i.inventorySrv.SaveCategory(&requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx, rp)
	}
}

func (i *InventoryHandler) UpdateCategory(ctx *gin.Context) {
	var (
		requestBody cmd.CategoryUpdateCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if err := i.inventorySrv.UpdateCategory(&requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx)
	}
}

func (i *InventoryHandler) RemoveCategoryByIds(ctx *gin.Context) {
	var (
		requestBody []string
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if err := i.inventorySrv.RemoveCategoryByIds(requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx)
	}
}

func (i *InventoryHandler) InStockOne(ctx *gin.Context) {
	var (
		requestBody cmd.StockSaveCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if string, err := i.inventorySrv.InStockOne(&requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx, string)
	}
}

func (i *InventoryHandler) QueryCategory(ctx *gin.Context) {
	var (
		requestBody query.CategoryPageQuery
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if data, err := i.inventorySrv.QueryCategory(requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx, data)
	}
}

func (i *InventoryHandler) QueryStock(ctx *gin.Context) {
	var (
		requestBody query.StockPageQuery
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		i.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if data, err := i.inventorySrv.QueryStock(requestBody); err != nil {
		i.SendFailure(ctx, err.Error())
	} else {
		i.SendSuccess(ctx, data)
	}
}
