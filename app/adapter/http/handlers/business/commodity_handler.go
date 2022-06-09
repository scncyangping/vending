package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/service"
	"vending/app/types/constants"
)

type CommodityHandler struct {
	*handlers.Handler
	commoditySrv service.CommoditySrv
}

func NewCommodityHandler(handler *handlers.Handler, commoditySrv service.CommoditySrv) *CommodityHandler {
	return &CommodityHandler{
		Handler:      handler,
		commoditySrv: commoditySrv,
	}
}

func (c *CommodityHandler) CreateCommodity(ctx *gin.Context) {
	var (
		requestBody cmd.CommoditySaveCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		c.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if rp, err := c.commoditySrv.Save(&requestBody); err != nil {
		c.SendFailure(ctx, err.Error())
	} else {
		c.SendSuccess(ctx, rp)
	}
}

func (c *CommodityHandler) UpdateCommodity(ctx *gin.Context) {
	var (
		updateBody cmd.CommodityUpdateCmd
	)
	err := ctx.ShouldBind(&updateBody)
	if err != nil {
		c.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if err := c.commoditySrv.Update(&updateBody); err != nil {
		c.SendFailure(ctx, err.Error())
	} else {
		c.SendSuccess(ctx)
	}
}

func (c *CommodityHandler) RemoveCommodity(ctx *gin.Context) {
	var (
		Ids struct {
			Ids []string `json:"ids"`
		}
	)
	err := ctx.ShouldBind(&Ids)
	if err != nil {
		c.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if err := c.commoditySrv.Delete(Ids.Ids); err != nil {
		c.SendFailure(ctx, err.Error())
	} else {
		c.SendSuccess(ctx)
	}
}

func (c *CommodityHandler) GetCommodity(ctx *gin.Context) {
	commodityId := ctx.Param("id")
	if commodityId == constants.EmptyStr {
		c.SendFailure(ctx, handlers.RequestParameterError, handlers.StatusText(handlers.RequestParameterError))
		return
	}
	if data := c.commoditySrv.Get(commodityId); data == nil {
		c.SendFailure(ctx, handlers.DataNotFound)
	} else {
		c.SendSuccess(ctx, data)
	}
}

func (c *CommodityHandler) UpCommodity(ctx *gin.Context) {
	commodityId := ctx.Param("id")
	if commodityId == constants.EmptyStr {
		c.SendFailure(ctx, handlers.RequestParameterError, handlers.StatusText(handlers.RequestParameterError))
		return
	}
	if data := c.commoditySrv.Up(commodityId); data == nil {
		c.SendFailure(ctx, handlers.DataNotFound)
	} else {
		c.SendSuccess(ctx, data)
	}
}

func (c *CommodityHandler) DownCommodity(ctx *gin.Context) {
	commodityId := ctx.Param("id")
	if commodityId == constants.EmptyStr {
		c.SendFailure(ctx, handlers.RequestParameterError, handlers.StatusText(handlers.RequestParameterError))
		return
	}
	if data := c.commoditySrv.Down(commodityId); data == nil {
		c.SendFailure(ctx, handlers.DataNotFound)
	} else {
		c.SendSuccess(ctx, data)
	}
}

func (c *CommodityHandler) ListCommodity(ctx *gin.Context) {
	var (
		query query.CommoditiesPageQuery
	)
	err := ctx.ShouldBind(&query)
	if err != nil {
		c.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if data, err := c.commoditySrv.Query(query); err != nil {
		c.SendFailure(ctx, handlers.DataNotFound)
	} else {
		c.SendSuccess(ctx, data)
	}
}
