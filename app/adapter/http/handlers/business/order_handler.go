package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/service"
	"vending/app/types/constants"
)

type OrderHandler struct {
	*handlers.Handler
	orderSrv service.OrderSrv
}

func NewOrderHandler(handler *handlers.Handler, orderSrv service.OrderSrv) *OrderHandler {
	return &OrderHandler{Handler: handler, orderSrv: orderSrv}
}

func (o *OrderHandler) CreateOrder(ctx *gin.Context) {
	var (
		requestBody cmd.CreateOrderCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if rp, err := o.orderSrv.CreateOrder(&requestBody); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx, rp)
	}
}

func (o *OrderHandler) OrderCallBack(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == constants.EmptyStr {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if err := o.orderSrv.OrderCallBack(orderId); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx)
	}
}

func (o *OrderHandler) Cancel(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == constants.EmptyStr {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if err := o.orderSrv.Cancel(orderId); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx)
	}
}

func (o *OrderHandler) GetTempOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == constants.EmptyStr {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if body, err := o.orderSrv.GetTempOrderById(orderId); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx, body)
	}
}

func (o *OrderHandler) GetOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == constants.EmptyStr {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if body, err := o.orderSrv.GetOrderById(orderId); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx, body)
	}
}

func (o *OrderHandler) Query(ctx *gin.Context) {
	var (
		requestBody query.OrderPageQuery
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		o.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

	if rp, err := o.orderSrv.Query(requestBody); err != nil {
		o.SendFailure(ctx, err.Error())
	} else {
		o.SendSuccess(ctx, rp)
	}
}
