//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"vending/app/adapter/http/server"
	service2 "vending/app/application/service"
	"vending/app/domain/aggregate/factory"
	"vending/app/domain/service"
	"vending/app/infrastructure/config"
	"vending/app/infrastructure/repository"
)

var providerSet = wire.NewSet(
	config.NewConfig,
	repository.NewRepository,
	factory.NewAggregate,
	service.NewService,
	service2.NewAppSrvManager,
	server.NewHandlers,
)

func NewHandler() *server.Handlers {
	panic(wire.Build(providerSet))
}
