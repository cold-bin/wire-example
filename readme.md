## introduce

本项目仅仅只是一个示例

使用wire实现依赖注入，避免层间依赖，既方便mock和测试，也方便维护

## 项目结构

```
.
├── cmd
│   └── di # injector
├── conf
├── dal
├── entity
├── handler
├── provider # di所需的provider
│   ├── kmq
│   ├── mysql
│   ├── redis
│   └── zap
├── router
└── service
```

## 食用指南

- 每一层都需要将依赖外部组件和下一层抽象到一个结构体里，然后每层（router,handler,service,dal）都按照业务拆分成实体并封装实体方法
- 这样所有的依赖都可以通过`wire`注入，而不是直接引用外部包的变量
- `wire`提供了一些很方便的api，帮助我们省掉工厂方法的创建，诸如`wire.Struct`等
- 但是有一些组件的对象，只有我们自己提供工厂方法来创建了，例如数据库句柄等

### injector

```go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"school-seckill-system/dal"
	"school-seckill-system/handler"
	"school-seckill-system/provider/kmq"
	"school-seckill-system/provider/mysql"
	"school-seckill-system/provider/redis"
	"school-seckill-system/provider/zap"
	"school-seckill-system/router"
	"school-seckill-system/service"
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
```

### provider

自提供工厂方法的provider示例：

```go
var (
	gdb  *gorm.DB
	once = &sync.Once{}
)

// 单例模式
func DBClient() *gorm.DB {
	if gdb != nil {
		return gdb
	}

	once.Do(func() {
		source := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.MysqlUsername,
			conf.MysqlPassword,
			conf.MysqlHost,
			conf.MysqlPort,
			conf.MysqlDbname,
		)
		db_, err := gorm.Open(mysql.Open(source), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		db, err := db_.DB()
		if err != nil {
			panic(err)
		}
		db.SetMaxOpenConns(conf.MysqlMaxOpenConns)
		db.SetMaxIdleConns(conf.MysqlMaxIdleConns)
		db.SetConnMaxLifetime(conf.MysqlConnMaxLifeMinutes * time.Minute)
		err = db_.Migrator().AutoMigrate()
		if err != nil {
			panic(err)
		}
		gdb = db_
	})
	return gdb
}

```