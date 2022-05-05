package mysql

import (
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

/**
 * @brief 加载数据库
 */

var (
	MysqlClient *gorm.DB
)

func Init() {
	dsn := conf.GlobalConfig.Mysql.Dsn
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.PrintAndDie(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		utils.PrintAndDie(err)
	}
	// conn_max_lifetime must be set, otherwise dbproxy will kill the conn 120s
	sqlDB.SetConnMaxLifetime(time.Duration(conf.GlobalConfig.Mysql.ConnMaxLifetime))
	sqlDB.SetMaxIdleConns(conf.GlobalConfig.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.GlobalConfig.Mysql.MaxOpenConns)
	MysqlClient = db
}
