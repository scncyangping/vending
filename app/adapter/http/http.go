package main

import (
	"github.com/gin-gonic/gin"
	"vending/app/ddd/adapter/http/sg"
)

func main() {
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
