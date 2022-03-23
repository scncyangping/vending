package log

import (
	"go.uber.org/zap"
	"vending/data/common/config"
	"vending/data/common/config/log/zap"
)

var ZapLogger *zap.SugaredLogger

// 创建zap logger
func NewLogger() {
	if ZapLogger == nil {
		ZapLogger = zapUtil.NewZapLogger(
			zapUtil.SetDevelopment(true),
			zapUtil.SetLevel(zap.DebugLevel),
			zapUtil.SetAppName(config.Base.Server.Name),
		).Sugar()
	}
}
