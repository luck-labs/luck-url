package env

import (
	"flag"
	"github.com/luck-labs/luck-url/common/consts/biz_const"
	"os"
)

var Env string

func init() {
	var defaultEnv = biz_const.EnvDevelopment
	if len(os.Getenv("LUCK_URL_ENV")) > 0 {
		defaultEnv = os.Getenv("LUCK_URL_ENV")
	}
	flag.StringVar(&Env, "LUCK_URL_ENV", defaultEnv, "set the environment")
}
