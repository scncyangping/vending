package main

import (
	"vending/app/adapter/http/routers"
	"vending/app/adapter/http/server"
	"vending/app/infrastructure/config"
	"vending/app/types/constants"
)

func run(mode string) {

}

func main() {
	config.NewConfig()
	h := NewHandler()

	server := server.NewHttpGin(constants.DebugMode)
	routers.InitRoute(server.Engine, h)
	server.Start()
}
