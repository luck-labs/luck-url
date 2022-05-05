package dao_mysql

import (
	"context"
	"github.com/luck-labs/luck-url/plugin/mysql"
	"gorm.io/gorm"
)

type BaseDao interface {
	GetDB(ctx context.Context) *gorm.DB
}

type baseDao struct {
	options
}

var baseDaoInstance = &baseDao{}

func NewBaseDao() BaseDao {
	return baseDaoInstance
}

// 选项模式
type options struct {
	OptionMaster     bool
	OptionForUpdate  bool
	OptionForceIndex string
	OptionShadow     bool
}

type Option func(opt *options)

/**
 * @brief 读主
 */
func SetMaster() Option {
	return func(opt *options) {
		opt.OptionMaster = true
	}
}

/**
 * @brief 间隙锁
 */
func SetForUpdate() Option {
	return func(opt *options) {
		opt.OptionForUpdate = true
	}
}

/**
 * @brief 指定索引
 */
func SetForceIndex(column string) Option {
	return func(opt *options) {
		opt.OptionForceIndex = column
	}
}

func (b *baseDao) GetDB(ctx context.Context) *gorm.DB {
	var resdb = getConn(ctx, false)
	return resdb
}
func (b *baseDao) GetDBFromMaster(ctx context.Context) *gorm.DB {
	return getConn(ctx, true)
}

func getConn(ctx context.Context, fromMaster bool) *gorm.DB {
	return mysql.MysqlClient.WithContext(ctx)
}
