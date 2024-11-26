package db

import (
	"context"
	cfg "demo1116/config"
	"demo1116/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rc *redis.Client

func initRedis() {
	// 提取配置
	c := cfg.Cfg.Redis
	Rc = redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.Db,
	})

	_, err := Rc.Ping(context.Background()).Result()
	if err != nil {
		logger.GLog.Info(fmt.Sprintf("Redis 连接失败 err = %v", err))
	} else {
		logger.GLog.Info(fmt.Sprintf("Redis 连接成功 ！！！"))
	}
}
