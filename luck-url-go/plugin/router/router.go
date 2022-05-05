package router

import (
	"github.com/luck-labs/luck-url/controller/short_url_api"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

/**
 * @brief 加载 httprouter
 */
var Handler http.Handler

func Init() {
	router := httprouter.New()
	// API
	router.POST("/v1/api/create", short_url_api.NewShortUrlController().Create)  // 创建短链
	router.GET("/v1/api/get/:suffix", short_url_api.NewShortUrlController().Get) // 解析短链
	router.GET("/v1/jump/:s", short_url_api.NewShortUrlController().Redirect)    // 重定向链接
	// cors跨域解决
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	Handler = handler
}
