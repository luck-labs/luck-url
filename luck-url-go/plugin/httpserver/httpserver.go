package httpserver

import (
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/router"
	"github.com/luck-labs/luck-url/plugin/utils"
	"golang.org/x/net/netutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/**
 * @brief 加载HttpServer
 */
func Init() {
	// HttpServer配置
	httpServerConfig := conf.GlobalConfig.HttpServer
	server := &http.Server{
		ReadTimeout: time.Second * 5,
		Addr:        httpServerConfig.Address,
		Handler:     router.Handler,
	}

	/**
	 * @brief 限流配置
	 * @doc   https://gist.github.com/caiofilipini/b85ad9f8db89eac9a6e9496e788c54da
	 */
	listener, err := net.Listen("tcp", httpServerConfig.Address)
	if err != nil {
		utils.PrintAndDie(err)
	}
	listener = netutil.LimitListener(listener, httpServerConfig.MaxConn)
	defer listener.Close()

	// HttpServer启动
	if err = server.Serve(listener); err != nil {
		utils.PrintAndDie(err)
	}
}
