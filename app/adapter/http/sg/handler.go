package sg

import (
	"go.uber.org/zap"
	"vending/app/domain/repo"
)

// Handler 具体业务服务聚合
type Handler struct {
	Logger *zap.SugaredLogger
}

type AuthHandler struct {
	Handler
	// user相关
	// jwt相关
}

func NewAuthHandler(logger *zap.SugaredLogger, repo repo.AuthServiceRepo) {

}
