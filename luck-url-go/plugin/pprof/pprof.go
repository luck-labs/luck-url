package pprof

import (
	"github.com/luck-labs/luck-url/plugin/conf"
	"net/http"
)

func Init() {
	pprofConfig := conf.GlobalConfig.Pprof
	go func() {
		_ = http.ListenAndServe(pprofConfig.Address, nil)
	}()
}
