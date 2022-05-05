package bloom

import (
	"context"
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/env"
	"github.com/luck-labs/luck-url/plugin/id"
	"github.com/luck-labs/luck-url/plugin/redis"
	"testing"
)

func TestInitBloom(t *testing.T) {
	ctx := context.Background()
	confPath := "../../conf/" + env.Env + "/app.toml" // 配置路径
	conf.Init(confPath)                               // Config 初始化
	redis.Init()                                      // Redis 初始化
	bloomFilter := NewDefaultBloomFilter("USER_BLOOM_TEST", "USER_BLOOM_TEST_CNT")
	uniqueStr := gocast.ToString(id.GetNextId(ctx))
	notSetStr := gocast.ToString(id.GetNextId(ctx) + 1)

	bloomFilter.AddString(ctx, gocast.ToString(uniqueStr))
	res := bloomFilter.TestString(ctx, gocast.ToString(uniqueStr))
	if res {
		fmt.Println("CONTAINS KEY")
	} else {
		fmt.Println("NOT CONTAINS KEY")
	}

	res = bloomFilter.TestString(ctx, gocast.ToString(notSetStr))
	if res {
		fmt.Println("CONTAINS KEY")
	} else {
		fmt.Println("NOT CONTAINS KEY")
	}
}
