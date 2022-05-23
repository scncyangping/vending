package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/handlers"
	"vending/app/application/dto"
	"vending/app/application/service/impl"
)

type AuthHandler struct {
	*handlers.Handler
	userSrv *impl.AuthSrvImp
}

// NewAuthHandler wire
func NewAuthHandler(auth *impl.AuthSrvImp) *AuthHandler {
	return &AuthHandler{
		Handler: handlers.NewHandler(),
		userSrv: auth,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var (
		requestBody dto.LoginRe
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
		re dto.RegisterRe
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
