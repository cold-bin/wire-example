//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"wire-example/dal"
	"wire-example/handler"
	"wire-example/provider/kmq"
	"wire-example/provider/mysql"
	"wire-example/provider/redis"
	"wire-example/provider/zap"
	"wire-example/router"
	"wire-example/service"
)

func BuildItemRouter() *router.ItemRouter {
	panic(
		wire.Build(
			wire.Struct(new(router.ItemRouter), "*"),
			wire.Struct(new(handler.ItemHandler), "*"),
			wire.Struct(new(service.ItemSvc), "*"),
			redis.RClient,
			kmq.KmqClient,
			zap.Logger,
			wire.Struct(new(dal.ItemDal), "*"),
			wire.Struct(new(dal.Base), "*"),
			mysql.DBClient,
		))
}
