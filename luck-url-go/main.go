package main

import (
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/env"
	"github.com/luck-labs/luck-url/plugin/httpserver"
	"github.com/luck-labs/luck-url/plugin/id"
	"github.com/luck-labs/luck-url/plugin/log"
	"github.com/luck-labs/luck-url/plugin/logo"
	"github.com/luck-labs/luck-url/plugin/pprof"
	"github.com/luck-labs/luck-url/plugin/redis"
	"github.com/luck-labs/luck-url/plugin/router"
	"github.com/luck-labs/luck-url/plugin/rpc"
)

func main() {
	confPath := "./conf/" + env.Env + "/app.toml" // 配置路径
	conf.Init(confPath)                           // Config 初始化
	log.Init()                                    // 日志初始化
	logo.Init()                                   // LOGO初始化
	id.Init()                                     // Snowflake初始化
	redis.Init()                                  // Redis 初始化
	router.Init()                                 // HttpRouter 初始化
	rpc.Init()                                    // RPC 初始化
	pprof.Init()                                  // Pprof 初始化
	httpserver.Init()                             // HttpServer 初始化
}
