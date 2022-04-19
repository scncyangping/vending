package main

import (
	"vending/app/adapter/http/sg"
	"vending/app/adapter/http/sg/routers"
	"vending/app/infrastructure/config"
	"vending/app/types/constants"
)

func NewHttp(path, mod string) {
	config.NewConfig(path)
	srv := sg.NewHttpGin(mod)
	routers.InitRoute(srv.Engine)
	srv.Start()
}

func main() {
	NewHttp("/Users/yapi/WorkSpace/VscodeWorkSpace/vending/cmd/config.yml", constants.DebugMode)
}
