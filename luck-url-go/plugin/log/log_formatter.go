package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/luck-labs/luck-url/common/consts/time_const"
	"github.com/sirupsen/logrus"
)

/**
 * @brief 自定义日志格式
 */

type CustomizedFormatter struct {
}

func (m *CustomizedFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}
	var logTime string // 日志时间
	var logMdu string  // 日志模块
	var logIdx string  // 日志标签
	logTime = entry.Time.Format(time_const.YYYY_MM_DD_HH_MM_SS)
	data := entry.Data
	//var ctx context.Context
	if _, ok := data["ctx"]; ok {
		//ctx = data["ctx"].(context.Context)
		delete(data, "ctx")
	}
	if _, ok := data["msg"]; !ok {
		if len(entry.Message) > 0 {
			data["msg"] = entry.Message
		}
	}
	if _, ok := data["mdu"]; ok {
		v := data["mdu"].(string)
		logMdu = v
		delete(data, "mdu")
	}
	if _, ok := data["idx"]; ok {
		v := data["idx"].(string)
		logIdx = v
		delete(data, "idx")
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	var logStr string
	logStr = fmt.Sprintf("[%s] [%s] [%s] [%s] %s\n", logTime, entry.Level, logMdu, logIdx, dataJson)
	buffer.WriteString(logStr)
	return buffer.Bytes(), nil
}
