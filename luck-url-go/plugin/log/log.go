package log

import (
	"context"
	"github.com/luck-labs/luck-url/common/consts/time_const"
	"github.com/luck-labs/luck-url/plugin/conf"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var LOG *logrus.Logger

func Init() {
	logConfig := conf.GlobalConfig.Log
	// INFO日志注册
	infoWriter, err := rotatelogs.New(
		logConfig.InfoLogPath+logConfig.LogFormat,
		rotatelogs.WithLinkName(logConfig.InfoLogPath),
		rotatelogs.WithMaxAge(time.Duration(time_const.Day*7)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(10)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	// ERROR日志注册
	errorWriter, err := rotatelogs.New(
		logConfig.ErrorLogPath+logConfig.LogFormat,
		rotatelogs.WithLinkName(logConfig.ErrorLogPath),
		rotatelogs.WithMaxAge(time.Duration(time_const.Day*7)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(10)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errorWriter,
		},
		&CustomizedFormatter{},
	))

	// logger.SetFormatter(&logrus.TextFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	logger.SetLevel(logrus.InfoLevel)
	LOG = logger
}

func Trace(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Trace()
}

func Debug(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Debug()
}

func Info(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Info()
}

func Warn(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Warn()
}

func Error(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Error()
}

func Fatal(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) {
	LOG.WithFields(getLogFields(ctx, logMdu, logIdx, logData)).Fatal()
}

func getLogFields(ctx context.Context, logMdu string, logIdx string, logData logrus.Fields) logrus.Fields {
	logData["ctx"] = ctx
	logData["mdu"] = logMdu
	logData["idx"] = logIdx
	return logData
}
