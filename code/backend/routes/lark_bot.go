package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/handlers/v1/admin"
	"github.com/tietiexx/bot/code/backend/utils"
	"net/http"
)

type LarkBotRoute struct{}

func (r *LarkBotRoute) Init(g *gin.RouterGroup) {
	var botRoutesInfo = constant.RoutesInfo{
		{
			Path:        "/lark_bots",
			Method:      http.MethodGet,
			HandlerFunc: admin.LarkBotList,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_bots",
			Method:      http.MethodPost,
			HandlerFunc: admin.LarkBotCreate,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_bots/:larkBotUUID",
			Method:      http.MethodGet,
			HandlerFunc: admin.LarkBotDetail,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_bots/:larkBotUUID",
			Method:      http.MethodPut,
			HandlerFunc: admin.LarkBotUpdate,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_bots/:larkBotUUID",
			Method:      http.MethodPatch,
			HandlerFunc: admin.LarkBotUpdatePartial,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_bots/:larkBotUUID",
			Method:      http.MethodDelete,
			HandlerFunc: admin.LarkBotDelete,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
	}

	// 遍历路由信息，注册路由
	utils.InitRoutesInfo(g, "/", botRoutesInfo)
}
