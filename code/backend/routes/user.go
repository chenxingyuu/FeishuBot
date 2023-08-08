package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/handlers/v1/admin"
	"github.com/tietiexx/bot/code/backend/utils"
	"net/http"
)

type UserRoute struct{}

func (r *UserRoute) Init(g *gin.RouterGroup) {
	var userRoutesInfo = constant.RoutesInfo{
		{
			Path:        "/login",
			Method:      http.MethodPost,
			HandlerFunc: admin.UserLogin,
		},
		{
			Path:        "/logout",
			Method:      http.MethodPost,
			HandlerFunc: admin.UserLogout,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
		{
			Path:        "/info",
			Method:      http.MethodGet,
			HandlerFunc: admin.UserInfo,
			RouteGroups: []constant.RouteGroup{jwtAuthGroup},
		},
	}
	// 遍历路由信息，注册路由
	utils.InitRoutesInfo(g, "/user", userRoutesInfo)
}
