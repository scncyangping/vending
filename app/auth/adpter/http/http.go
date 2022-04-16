package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"vending/app/auth/adpter/http/auth_handles"
	"vending/app/auth/adpter/http/routers"
	"vending/app/auth/domain/service"
	"vending/app/auth/infrastructure/pkg/log"
)

func NewHttp() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	h := auth_handles.NewHandles(s, auth)
	routers.SetRouters(g, h)
	server := &http.Server{
		Addr:           s.NetConf.ServerAddr,
		Handler:        g,
		ReadTimeout:    time.Duration(s.NetConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(s.NetConf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.GetLogger().Info("auth server start success", zap.Any("addr", s.NetConf.ServerAddr))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err.(interface{}))
		}
	}()
}
