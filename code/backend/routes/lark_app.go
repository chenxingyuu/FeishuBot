package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/handlers/v1/admin"
	"github.com/tietiexx/bot/code/backend/utils"
	"net/http"
)

type LarkAppRoute struct{}

func (r *LarkAppRoute) Init(g *gin.RouterGroup) {
	var botRoutesInfo = constant.RoutesInfo{
		{
			Path:        "/lark_apps",
			Method:      http.MethodGet,
			HandlerFunc: admin.LarkAppList,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_apps",
			Method:      http.MethodPost,
			HandlerFunc: admin.LarkAppCreate,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_apps/:larkAppUUID",
			Method:      http.MethodGet,
			HandlerFunc: admin.LarkAppDetail,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_apps/:larkAppUUID",
			Method:      http.MethodPut,
			HandlerFunc: admin.LarkAppUpdate,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_apps/:larkAppUUID",
			Method:      http.MethodPatch,
			HandlerFunc: admin.LarkAppUpdatePartial,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/lark_apps/:larkAppUUID",
			Method:      http.MethodDelete,
			HandlerFunc: admin.LarkAppDelete,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
	}

	// 遍历路由信息，注册路由
	utils.InitRoutesInfo(g, "/", botRoutesInfo)
}
