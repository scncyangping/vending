package sg

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type ServerConf struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}
type HttpGin struct {
	Conf   *ServerConf
	Engine *gin.Engine
	Logger *zap.SugaredLogger
}

type HttpRouteGroup struct {
	GroupPath string
	Groups    []func(*gin.RouterGroup)
}

func NewHttpRouteGroup(path string, gs ...func(*gin.RouterGroup)) *HttpRouteGroup {
	if path == "" {
		panic("routers path is empty")
	}
	gps := make([]func(*gin.RouterGroup), 0)
	if len(gs) > 0 {
		for _, v := range gs {
			gps = append(gps, v)
		}
	}
	return &HttpRouteGroup{
		GroupPath: path,
		Groups:    gps,
	}
}

func (e *HttpGin) NewHttpGin(mod string, hrg ...*HttpRouteGroup) *HttpGin {
	gin.SetMode(mod)
	g := gin.Default()

	for _, v := range hrg {
		gr := e.Engine.Group(v.GroupPath)
		for _, k := range v.Groups {
			k(gr)
		}
	}
	return &HttpGin{Engine: g}
}

func (e *HttpGin) Start() {
	server := &http.Server{
		Addr:           e.Conf.Addr,
		Handler:        e.Engine,
		ReadTimeout:    time.Duration(e.Conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(e.Conf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	e.Logger.Info("Vending Server Start Success", zap.Any("addr", e.Conf.Addr))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			e.Logger.Error("Vending Server Error!")
			panic(err)
		}
	}()
}
