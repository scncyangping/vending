package auth_handles

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vending/app/auth/domain/service"
	"vending/app/auth/infrastructure/conf"
	"vending/app/auth/infrastructure/pkg/hcode"
	"vending/app/auth/infrastructure/pkg/log"
)

type Handles struct {
	config *conf.AppConfig
	auth   service.AuthSrv
}

func NewHandles(conf *conf.AppConfig, auth service.AuthSrv) *Handles {
	return &Handles{
		config: conf,
		auth:   auth,
	}
}
func (h *Handles) GetUid(g *gin.Context) int {
	key, ok := g.Get("uid")
	if !ok {
		h.ResponseErr(g, hcode.ParameterErr)
		return 0
	}
	return key.(int)
}

func (h *Handles) ResponseErr(g *gin.Context, err error) {
	code := hcode.Cause(err)
	data := gin.H{
		"code": code.Code(),
		"data": "",
		"msg":  code.Message(g.GetHeader("lang")),
	}
	log.GetLogger().Debug("ResponseErr", zap.Any("res", data))
	g.JSON(200, data)
}
func (h *Handles) ResponseSuccess(g *gin.Context) {
	info := gin.H{
		"code": hcode.OK,
		"data": "",
		"msg":  hcode.OK.Message(g.GetHeader("lang")),
	}
	log.GetLogger().Debug("ResponseData", zap.Any("res", info))
	g.JSON(200, info)
}

func (h *Handles) ResponseData(g *gin.Context, data interface{}) {
	info := gin.H{
		"code": hcode.OK,
		"data": data,
		"msg":  hcode.OK.Message(g.GetHeader("lang")),
	}
	log.GetLogger().Debug("ResponseData", zap.Any("res", info))
	g.JSON(200, info)
}
