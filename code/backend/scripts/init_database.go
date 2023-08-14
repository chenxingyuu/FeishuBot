package main

import (
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/global"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils/passwordutil"
)

func main() {
	// 加载配置文件
	global.InitViper()
	// 初始化数据库连接
	global.InitMySQL(global.MySQLConf)

	// 同步数据库结构
	_ = global.MySQLClient.AutoMigrate(
		&database.User{},
		&database.LarkApp{},
		&database.LarkBot{},
		&database.WebhookTask{},
		&database.Receiver{},
	)

	// 初始化静态数据
	if _, err := services.UserByUsername("admin"); err != nil {
		hashedPassword, _ := passwordutil.EncryptPassword("123456")
		user := database.User{Username: "admin", HashedPassword: hashedPassword}
		_ = global.MySQLClient.Create(&user)
	}

}
