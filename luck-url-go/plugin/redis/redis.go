package redis

import (
	"context"
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/utils"
	"github.com/go-redis/redis/v8"
	"time"
)

/**
 * @brief 加载Redis缓存
 */

var (
	RedisClient *redis.Client
)

func Init() {
	ctx := context.Background()
	redisConf := conf.GlobalConfig.Redis
	connTimeout := time.Duration(redisConf.ConnTimeOutMS) * time.Second
	readTimeout := time.Duration(redisConf.ReadTimeOut) * time.Second
	writeTimeout := time.Duration(redisConf.WriteTimeOut) * time.Second
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisConf.Address[0],
		Password:     redisConf.Auth,     // no password set
		DB:           0,                  // use default DB
		PoolSize:     redisConf.PoolSize, // pool size
		DialTimeout:  connTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		utils.PrintAndDie(err)
	}
}
