package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils/jwtutil"
	"github.com/tietiexx/bot/code/backend/utils/response"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.UnauthorizedJsonResponse(c, "Missing token")
			c.Abort()
			return
		}

		// 检查 Token 格式是否为 "Bearer <token>"
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.UnauthorizedJsonResponse(c, "Invalid token format")
			c.Abort()
			return
		}

		tokenString := headerParts[1]

		claims, err := jwtutil.VerifyAccessToken(tokenString)
		if err != nil {
			response.UnauthorizedJsonResponse(c, "Invalid token")
			c.Abort()
			return
		}
		// 校验 user
		user, err := services.UserById(claims.UserID)
		if err != nil {
			response.UnauthorizedJsonResponse(c, "Inviable token")
			c.Abort()
			return
		}
		// 在上下文中设置用户信息
		c.Set("userID", claims.UserID)
		c.Set("userInfo", user)
		c.Next()
	}
}
