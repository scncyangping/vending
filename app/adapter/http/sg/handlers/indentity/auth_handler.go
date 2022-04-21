package indentity

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/sg/handlers"
	"vending/app/domain/dto"
	"vending/app/domain/service"
	authS "vending/app/domain/service/auth"
)

type AuthHandler struct {
	*handlers.Handler
	AuthSrv service.AuthSrv
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		Handler: handlers.NewHandler(),
		AuthSrv: authS.NewJwtTokenAuth(),
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var (
		requestBody dto.AuthenticationRe
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		h.SendFailure(ctx, err.Error())
		return
	}

	if requestBody.UserName == "aw" && requestBody.Pwd == "123" {
		// 创建token
		if token, err := h.AuthSrv.GenerateToken(&dto.JwtAuthTokenRe{UserName: requestBody.UserName}); err != nil {
			h.SendFailure(ctx, "jwt generate error")
			return
		} else {
			h.SendSuccess(ctx, token)
		}
	} else {
		h.SendFailure(ctx, "password or name error")
	}
}

func (h *AuthHandler) Who(ctx *gin.Context) {
	h.SendSuccess(ctx, "I am A W")
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var (
		requestBody dto.UserRegisterRq
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		h.SendFailure(ctx, handlers.ParameterConvertError, handlers.StatusText(handlers.ParameterConvertError))
		return
	}

}
