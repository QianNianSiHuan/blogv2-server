package core

import (
	"blogv2/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	ProgressbarMsg <- "redis初始化..."
	r := global.Config.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("redis连接成功")
	return redisDB
}
