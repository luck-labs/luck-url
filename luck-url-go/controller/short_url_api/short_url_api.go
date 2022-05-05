package short_url_api

import (
	"github.com/luck-labs/luck-url/common/consts/retcode"
	"github.com/luck-labs/luck-url/common/dto/http_dto"
	"github.com/luck-labs/luck-url/controller"
	"github.com/luck-labs/luck-url/plugin/log"
	"github.com/luck-labs/luck-url/service/rpcservice"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

/**
 * @brief 短链接
 */

type ShortUrlController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)   // 创建短链
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)      // 解析短链
	Redirect(w http.ResponseWriter, r *http.Request, params httprouter.Params) // 重定向
}

type shortUrlController struct {
	controller.Base
}

var shortUrlControllerInstance = &shortUrlController{}

func NewShortUrlController() ShortUrlController {
	return shortUrlControllerInstance
}

func (ctrl *shortUrlController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	log.Info(ctx, log.MduController, log.IdxControllerShortUrl, logrus.Fields{})

	var req http_dto.UrlCreateReqDto
	if err := ctrl.GetParams(r, &req); err != nil {
		ctrl.EchoRet(w, r, retcode.ErrParameterWrong, err.Error(), nil)
		return
	}

	ret, rsp, err := rpcservice.CreateUrl(ctx, req)
	if err != nil {
		ctrl.EchoRet(w, r, ret, err.Error(), rsp)
		return
	}

	ctrl.EchoRet(w, r, retcode.Success, "SUCCESS", rsp)
	return
}

func (ctrl *shortUrlController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	log.Info(ctx, log.MduController, log.IdxControllerShortUrl, logrus.Fields{})
	/*
		var req http_dto.UrlGetReqDto
		if err := ctrl.GetParams(r, &req); err != nil {
			ctrl.EchoRet(w, r, retcode.ErrParameterWrong, err.Error(), nil)
			return
		}
	*/

	ret, rsp, err := rpcservice.GetUrl(ctx, http_dto.UrlGetReqDto{ShortUrlSuffix: params.ByName("suffix")})
	if err != nil {
		ctrl.EchoRet(w, r, ret, err.Error(), rsp)
		return
	}

	ctrl.EchoRet(w, r, retcode.Success, "SUCCESS", rsp)
	return
}

func (ctrl *shortUrlController) Redirect(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	log.Info(ctx, log.MduController, log.IdxControllerShortUrl, logrus.Fields{})

	ret, rsp, err := rpcservice.GetUrl(ctx, http_dto.UrlGetReqDto{ShortUrlSuffix: params.ByName("s")})
	if err != nil {
		ctrl.EchoRet(w, r, ret, err.Error(), rsp)
		return
	}

	http.Redirect(w, r, rsp.Url, http.StatusFound)
	return
}
