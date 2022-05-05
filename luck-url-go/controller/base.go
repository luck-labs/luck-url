package controller

/**
 * @brief Controller基础
 */

import (
	"github.com/luck-labs/luck-url/common/consts/retcode"
	"github.com/luck-labs/luck-url/plugin/log"
	"github.com/go-playground/validator/v10"
	"github.com/liamylian/jsontime"
	"github.com/ppltools/binding"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"strings"
)

type Base struct {
	validate *validator.Validate
}

var baseInstance = &Base{
	validate: validator.New(),
}

func NewBaseController() *Base {
	return baseInstance
}

/**
 * @brief 获取参数
 */
func (b *Base) GetParams(r *http.Request, params interface{}) (err error) {
	/**
	 * @brief    r.request 转 Struct
	 * @document github.com/ppltools/binding
	 * 			 github.com/gin-gonic/gin/binding
	 */
	err = binding.Default(r.Method, r.Header.Get("Content-Type")).Bind(r, params)
	if err != nil {
		return err
	}

	/**
	 * @brief   参数validate
	 * @document https: //github.com/go-playground/validator
	 */
	if err := baseInstance.validate.Struct(params); err != nil {
		return err
	}
	return err
}

func (bs *Base) EchoRet(w http.ResponseWriter, r *http.Request, errno retcode.RetCode, errmsg string, data interface{}) (retcode.RetCode, string) {
	bs.EchoJSON(w, r, map[string]interface{}{
		"errno":  errno,
		"errmsg": errmsg,
		"data":   data,
	})
	return errno, errmsg
}

// EchoJSON json格式输出
func (bs *Base) EchoJSON(w http.ResponseWriter, r *http.Request, body interface{}) {
	//add jsontime to support timeformat
	json := jsontime.ConfigWithCustomTimeFormat

	if cType := w.Header().Get("Content-Type"); cType == "" {
		w.Header().Set("Content-Type", "application/json")
	}
	b, err := json.Marshal(body)
	if err != nil {
		bs.echo(w, r, []byte(`{"errno":1, "errmsg":"`+err.Error()+`"}`))
	} else {
		bs.echo(w, r, b)
	}
}

// Echo 原始输出,包含tracelog
func (bs *Base) echo(w http.ResponseWriter, req *http.Request, body []byte) {
	if cType := w.Header().Get("Content-Type"); cType == "" {
		w.Header().Set("Content-Type", "text/plain")
	}
	w.Write(body)
}

func (bs *Base) ExceptionHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		stackInfo := strings.Replace(string(debug.Stack()), "\n", "", 1)
		log.Error(r.Context(), log.MduRecovery, log.IdxControllerExceptionHandler, logrus.Fields{
			"tag":       "ExceptionHandlerFunc",
			"err":       err,
			"stackInfo": stackInfo,
		})
		bs.EchoRet(w, r, http.StatusInternalServerError, "系统错误", nil)
	}
}
