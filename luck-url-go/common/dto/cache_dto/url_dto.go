package cache_dto

type UrlData struct {
	OriginalUrl    string `json:"original_url"`     // 原始URL
	ShortUrlPrefix string `json:"short_url_prefix"` // 短链接前缀
}
