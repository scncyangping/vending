package adpter

import (
	"go.uber.org/zap"
	"vending/app/auth/adpter/http"
	"vending/app/auth/domain/service"
	"vending/app/auth/infrastructure/conf"
	"vending/app/auth/infrastructure/pkg/log"
	"vending/common/util"
)

type Server struct {
	conf *conf.AppConfig
	log  *log.Logger
	auth service.AuthSrv
}

func NewSrv(c *conf.AppConfig, log *log.Logger, auth service.AuthSrv) *Server {
	s := &Server{conf: c, log: log, auth: auth}
	s.Init()
	return s
}
func (s *Server) Init() {
	// hcode.Click()
}

func (s *Server) RunApp() {
	http.NewHttp(s.conf, s.auth)
	util.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("auth server exit", zap.Any("addr", s.conf.NetConf.ServerAddr))
	})
}

func (s *Server) Close() {

}
