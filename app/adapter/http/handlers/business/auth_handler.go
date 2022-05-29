package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/application/cqe/cmd"
	"vending/app/application/service"
)

type AuthHandler struct {
	*handlers.Handler
	userSrv service.AuthSrv
}

func NewAuthHandler(srv service.AuthSrv) *AuthHandler {
	return &AuthHandler{
		Handler: handlers.NewHandler(),
		userSrv: srv,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var (
		requestBody cmd.LoginCmd
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		h.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if rp, err := h.userSrv.Login(&requestBody); err != nil {
		h.SendFailure(ctx, err.Error())
	} else {
		h.SendSuccess(ctx, rp)
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var (
		re cmd.RegisterCmd
	)
	err := ctx.ShouldBind(&re)
	if err != nil {
		h.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}
	if uId, err := h.userSrv.Register(&re); err != nil {
		h.SendFailure(ctx, err.Error())
	} else {
		h.SendSuccess(ctx, uId)
	}
}
