package retcode

/**
 * @brief 返回值
 */

type RetCode int

const (
	Success           RetCode = 0    // 成功
	Failed            RetCode = 500  // 系统异常
	ErrJsonMarshal    RetCode = 1000 // JSON格式化异常
	ErrJsonUnmarshal  RetCode = 1001 // JSON格式化异常
	ErrParameterWrong RetCode = 1002 // 参数相关 参数错误或缺失
)

const (
	ErrMySQLWriteFail RetCode = 2001 // 数据库写入失败
	ErrMySQLReadFail  RetCode = 2002 // 数据库读取失败
	ErrRedisWriteFail RetCode = 2101 // 缓存写入失败
	ErrRedisReadFail  RetCode = 2102 // 缓存读取失败
)

const (
	ErrHitBloomUserNotFound RetCode = 10001 // 布隆过滤器
)
