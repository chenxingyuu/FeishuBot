package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/middlewares"
)

var jwtAuthGroup = constant.RouteGroup{
	MiddleWares: []gin.HandlerFunc{middlewares.JWTAuthMiddleware()},
}
