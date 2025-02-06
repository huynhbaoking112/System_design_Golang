package initialize

import (
	"context"
	"fmt"

	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", global.Config.RedisSetting.Host, global.Config.RedisSetting.Port),
		Password: global.Config.RedisSetting.Password, // no password set
		DB:       global.Config.RedisSetting.Database, // use default DB
		PoolSize: global.Config.RedisSetting.Poolsize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Initializing Redis fail", zap.Error(err))
	} else {
		global.Rdb = rdb
		global.Logger.Info("Initializing Redis Successfully")
	}
}
