// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func NewHandler() *server.Handlers {
	repositoryRepository := repository.NewRepository()
	serviceService := service.NewService(repositoryRepository)
	appSrvManager := service2.NewAppSrvManager(serviceService)
	handlers := server.NewHandlers(appSrvManager)
	return handlers
}

// wire.go:

var providerSet = wire.NewSet(config.NewConfig, repository.NewRepository, factory.NewAggregate, service.NewService, service2.NewAppSrvManager, server.NewHandlers)
