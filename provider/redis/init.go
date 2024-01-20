package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"wire-example/conf"
)

var (
	rdb  *redis.Client
	once = &sync.Once{}
)

// 单例模式
func RClient() *redis.Client {
	if rdb != nil {
		return rdb
	}

	once.Do(func() {
		c := redis.NewClient(&redis.Options{
			Addr:           conf.RedisAddr,
			MaxIdleConns:   conf.RedisMaxIdle,
			MaxActiveConns: conf.RedisMaxActive,
			Username:       conf.RedisUser,
			Password:       conf.RedisPassword,
		})

		if err := rdb.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
		rdb = c
	})

	return rdb
}
