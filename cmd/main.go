package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/global"
	"github.com/tietiexx/bot/code/backend/middlewares"
	"github.com/tietiexx/bot/code/backend/routes"
)

func SetupApp() {
	// 加载配置文件
	global.InitViper()
	// 初始化 MySQL
	global.InitMySQL(global.MySQLConf)
	// 初始化 Redis
	global.InitRedis(global.RedisConf)
}

func StartApp() {
	app := gin.New()
	// 初始化全局中间件
	app.Use(middlewares.LogMiddleware())
	app.Use(middlewares.GinCustomLogger())
	app.Use(middlewares.CorsMiddleware())
	// 初始化路由
	routes.Init(app)
	// 程序运行
	err := app.Run(global.BotAdminConf.Addr())
	if err != nil {
		return
	}

}

func main() {
	SetupApp()
	StartApp()
}
