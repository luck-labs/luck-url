package log

const (
	MduUndefined    = "mdu_undefined"
	MduMySQL        = "mdu_mysql"
	MduRedis        = "mdu_redis"
	MduController   = "mdu_controller"
	MduRpcFramework = "mdu_rpc_framework"
	MduRpcService   = "mdu_rpc_service"
	MduRecovery     = "mdu_recovery"
	MduRedisBloom   = "mdu_redis_bloom"
	MduSnowflake    = "mdu_snowflake"
	MduTest         = "mdu_test"
)

const (
	IdxUndefined                  = "idx_undefined"
	IdxControllerUser             = "idx_controller_user"
	IdxControllerShortUrl         = "idx_controller_short_url"
	IdxControllerExceptionHandler = "idx_controller_exception_handler"
	IdxRpcUser                    = "idx_rpc_user"
	IdxModelUser                  = "idx_model_user"
	IdxModelCacheUser             = "idx_model_cache_user"
	IdxRedisBloomTest             = "idx_redis_bloom_test"
	IdxRedisBloomAdd              = "idx_redis_bloom_add"
	IdxRedisBloomClear            = "idx_redis_bloom_clear"
	IdxRpcFrameworkClient         = "idx_rpc_framework_client"
	IdxSnowflakeGenId             = "idx_snowflake_genid"
)
