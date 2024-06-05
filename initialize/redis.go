package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-web-new/global"
	"go-web-new/utils"
)

func inItRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort),
		Password: utils.RedisPassword,
		DB:       3,
	})
	_, err := global.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("redis ping error")
	}

}
