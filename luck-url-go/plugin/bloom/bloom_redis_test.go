package bloom

import (
	"context"
	"fmt"
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/env"
	"github.com/luck-labs/luck-url/plugin/redis"
	"testing"
)

/**
 * @brief 写入
 */
func TestMSetBit(t *testing.T) {
	ctx := context.Background()
	confPath := "../../conf/" + env.Env + "/app.toml" // 配置路径
	conf.Init(confPath)                               // Config 初始化
	redis.Init()                                      // Redis 初始化
	key := "BLOOM_TEST"
	locations := make([]uint32, 0)
	locations = append(locations, uint32(100001))
	res, err := MSetBit(ctx, key, locations)
	if err != nil {
		fmt.Println(res)
	}
}

/**
 * @brief 读取
 */
func TestMGetBit(t *testing.T) {
	ctx := context.Background()
	confPath := "../../conf/" + env.Env + "/app.toml" // 配置路径
	conf.Init(confPath)                               // Config 初始化
	redis.Init()                                      // Redis 初始化
	key := "BLOOM_TEST"
	locations := make([]uint32, 0)
	locations = append(locations, uint32(100001))
	res, err := MGetBit(ctx, key, locations)
	if err != nil {
		fmt.Println(res)
	}
}
