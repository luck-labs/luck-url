package log

import (
	"context"
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/env"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	ctx := context.Background()
	confPath := "../../conf/" + env.Env + "/app.toml" // 配置路径
	conf.Init(confPath)                               // Config 初始化
	Init()
	Info(ctx, MduTest, IdxUndefined, logrus.Fields{
		"user_name": "alexhan",
	})
	Error(ctx, MduTest, IdxUndefined, logrus.Fields{
		"errmsg": "db error",
	})
	Debug(ctx, MduTest, IdxUndefined, logrus.Fields{
		"debugmsg": "debugging",
	})
}
