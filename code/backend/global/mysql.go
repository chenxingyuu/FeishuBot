package global

import (
	"fmt"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/utils/customlogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	mysqlOnce   sync.Once
	MySQLClient *gorm.DB
)

func mysqlConn(config *constant.MySQLConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	mysqlOnce.Do(
		func() {
			// 初始化 MySQL 连接
			db, err = gorm.Open(mysql.New(mysql.Config{
				DSN:                       config.DSN(), // DSN data source name
				DefaultStringSize:         256,          // string 类型字段的默认长度
				DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
				DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
				DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
				SkipInitializeWithVersion: false,        // 根据当前 MySQL 版本自动配置
			}), &gorm.Config{
				Logger: &customlogger.GormCustomLogger{
					GormLogger: logger.Default.LogMode(logger.Info),
				},
			})
		})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}
	return db, err
}

func InitMySQL(conf *constant.MySQLConfig) {
	conn, err := mysqlConn(conf)
	if err != nil {
		panic(err)
		return
	}
	MySQLClient = conn
}
