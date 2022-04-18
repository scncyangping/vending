package http

import (
	"github.com/gin-gonic/gin"
	"vending/app/adapter/http/sg"
)

func NewHttp() {
	httpGin := &sg.HttpGin{
		Conf: &sg.ServerConf{
			Addr:         "localhost:9079",
			ReadTimeout:  10000,
			WriteTimeout: 10000,
		},
	}
	groups := []*sg.HttpRouteGroup{
		sg.NewHttpRouteGroup("v1"),
	}

	httpGin.NewHttpGin(gin.DebugMode, groups...).Start()
}
