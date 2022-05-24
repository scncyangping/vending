package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vending/app/adapter/http/handlers/business"
	"vending/app/application/service"
	"vending/app/infrastructure/config"
	"vending/app/infrastructure/pkg/log"
)

type Handlers struct {
	AuthHandler *business.AuthHandler
}

// NewHandlers wire
func NewHandlers(srvM *service.AppSrvManager) *Handlers {
	return &Handlers{
		AuthHandler: business.NewAuthHandler(srvM.AuthSrv),
	}
}

type HttpGin struct {
	Conf   *config.Server
	Engine *gin.Engine
	Logger *zap.SugaredLogger
}

func NewHttpGin(mod string) *HttpGin {
	gin.SetMode(mod)
	g := gin.Default()
	return &HttpGin{Engine: g, Conf: config.Base.Server, Logger: log.Logger()}
}

func (e *HttpGin) Start() {
	server := &http.Server{
		Addr:           e.Conf.Addr,
		Handler:        e.Engine,
		ReadTimeout:    time.Duration(e.Conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(e.Conf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	e.Logger.Infof("Vending Server Start Success %s", e.Conf.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			e.Logger.Errorf("Vending Server Error! %s", e.Conf.Addr)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	e.Logger.Errorf("Shutdown Server ...%s", e.Conf.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		e.Logger.Errorf("Server Shutdown: %v %v \n", e.Conf.Addr, err)
	}
	select {
	case <-ctx.Done():
		e.Logger.Errorf("timeout of 10 seconds. %s", e.Conf.Addr)
	}

	e.Logger.Errorf("Server exiting, %s", e.Conf.Addr)
}
