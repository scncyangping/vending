//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"vending/app/auth/adpter"
	"vending/app/auth/domain/aggregate"
	"vending/app/auth/domain/service"
	"vending/app/auth/infrastructure/conf"
	"vending/app/auth/infrastructure/pkg/database/mongo"
	"vending/app/auth/infrastructure/pkg/database/redis"
	"vending/app/auth/infrastructure/pkg/log"
	"vending/app/auth/infrastructure/repository"
)

//go:generate wire
var providerSet = wire.NewSet(
	conf.NewViper,
	conf.NewAppConfigCfg,
	conf.NewLoggerCfg,
	conf.NewRedisConfig,
	conf.NewMongoConfig,
	log.NewLogger,
	redis.NewRedis,
	mongo.NewMongo,
	repository.NewRepository,
	aggregate.NewFactory,
	service.NewService,
	adpter.NewSrv,
)

func NewApp() (*adpter.Server, error) {
	panic(wire.Build(providerSet))
}
