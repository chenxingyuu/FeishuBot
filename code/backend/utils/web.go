package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/database"
)

func ContextGetUser(c *gin.Context) (user *database.User) {
	if val, ok := c.Get("userInfo"); ok && val != nil {
		user, _ = val.(*database.User)
	}
	return
}

func InitRoutesInfo(g *gin.RouterGroup, relativePath string, routesInfo constant.RoutesInfo) {
	Router := g.Group(relativePath)
	for _, route := range routesInfo {
		var handlers []gin.HandlerFunc
		for _, group := range route.RouteGroups {
			handlers = append(handlers, group.MiddleWares...)
		}
		handlers = append(handlers, route.HandlerFunc)
		Router.Handle(route.Method, route.Path, handlers...)
	}
}
