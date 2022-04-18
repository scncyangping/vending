package sg

import "go.uber.org/zap"

// Handler 具体业务服务聚合
type Handler struct {
	Logger *zap.SugaredLogger
}

type AuthHandler struct {
	Handler
	// user相关
	// jwt相关
}
