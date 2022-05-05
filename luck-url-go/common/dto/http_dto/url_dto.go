package http_dto

// 短链创建
type UrlCreateReqDto struct {
	Url            string `json:"url" validate:"required"`              // 链接
	ShortUrlPrefix string `json:"short_url_prefix" validate:"required"` // 短链接前缀
}

type UrlCreateRspDto struct {
	ShortUrl string `json:"short_url" validate:"required"` // 短链接
}

// 解析短链
type UrlGetReqDto struct {
	ShortUrlSuffix string `json:"short_url_suffix" validate:"required"` // 短链接
}

type UrlGetRspDto struct {
	Url string `json:"url" validate:"required"` // 长链接
}
