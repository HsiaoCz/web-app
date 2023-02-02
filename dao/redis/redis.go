package redis

import (
	"fmt"
	"web_app/settings"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(rcg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rcg.Host, rcg.Port),
		Password: rcg.Password,
		DB:       rcg.DB,
		PoolSize: rcg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis ping err:", zap.Error(err))
		return
	}
	return err
}

func Close() {
	_ = rdb.Close()
}
