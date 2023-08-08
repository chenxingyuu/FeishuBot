package constant

import "github.com/gin-gonic/gin"

type BaseRoute interface {
	Init(g *gin.RouterGroup)
}

type RouteGroup struct {
	MiddleWares []gin.HandlerFunc
}

type RouteInfo struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
	RouteGroups []RouteGroup
}

type RoutesInfo []RouteInfo
