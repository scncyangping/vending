//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"vending/app/adapter/http/handlers/business"
	"vending/app/adapter/http/server"
	"vending/app/application/service/impl"
	"vending/app/domain/aggregate"
	"vending/app/domain/service"
	"vending/app/infrastructure/config"
	"vending/app/infrastructure/repository"
)

var providerSet = wire.NewSet(
	config.NewConfig,
	repository.NewRepository,
	aggregate.NewAggregate,
	service.NewService,
	impl.NewAuthSrvImp,
	business.NewAuthHandler,
	server.NewHandlers,
)

func NewHandler() *server.Handlers {
	panic(wire.Build(providerSet))
}
