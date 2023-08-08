package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
)

func Init(e *gin.Engine) {
	routerGroup := e.Group("/api/v1")

	var userRoute constant.BaseRoute = &UserRoute{}
	userRoute.Init(routerGroup)

	var larkAppRoute constant.BaseRoute = &LarkAppRoute{}
	larkAppRoute.Init(routerGroup)

	var larkBotRoute constant.BaseRoute = &LarkBotRoute{}
	larkBotRoute.Init(routerGroup)
}
