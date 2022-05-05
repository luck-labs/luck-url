package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/luck-labs/luck-url/common/consts/retcode"
	"github.com/luck-labs/luck-url/common/dto/cache_dto"
	"github.com/luck-labs/luck-url/common/dto/http_dto"
	"github.com/luck-labs/luck-url/plugin/hash"
	"github.com/luck-labs/luck-url/plugin/redis"
	"strconv"
)

/**
 * @brief 创建短链
 */
func CreateUrl(ctx context.Context, req http_dto.UrlCreateReqDto) (retcode.RetCode, http_dto.UrlCreateRspDto, error) {
	suffix, hashid := hash.EncodeHashStr()
	jsonData, err := json.Marshal(&cache_dto.UrlData{
		OriginalUrl:    req.Url,
		ShortUrlPrefix: req.ShortUrlPrefix,
	})
	if err != nil {
		return retcode.ErrJsonUnmarshal, http_dto.UrlCreateRspDto{}, err
	}
	_, err = redis.RedisClient.Set(ctx, strconv.FormatInt(hashid, 10), string(jsonData), 0).Result()
	if err != nil {
		return retcode.ErrRedisWriteFail, http_dto.UrlCreateRspDto{}, err
	}
	return retcode.Success, http_dto.UrlCreateRspDto{
		ShortUrl: fmt.Sprintf("%s/%s", req.ShortUrlPrefix, suffix),
	}, nil
}

/**
 * @brief 解析短链
 */
func GetUrl(ctx context.Context, req http_dto.UrlGetReqDto) (retcode.RetCode, http_dto.UrlGetRspDto, error) {
	suffix := req.ShortUrlSuffix
	hashid := hash.DecodeHashStr(suffix)
	jsonData, err := redis.RedisClient.Get(ctx, strconv.FormatInt(hashid, 10)).Result()
	if err != nil {
		return retcode.ErrRedisReadFail, http_dto.UrlGetRspDto{}, err
	}
	var urlDto cache_dto.UrlData
	if err := json.Unmarshal([]byte(jsonData), &urlDto); err != nil {
		return retcode.ErrJsonUnmarshal, http_dto.UrlGetRspDto{}, err
	}
	return retcode.Success, http_dto.UrlGetRspDto{
		Url: urlDto.OriginalUrl,
	}, nil
}
