package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/converters"
	"github.com/tietiexx/bot/code/backend/models"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils"
	"github.com/tietiexx/bot/code/backend/utils/jwtutil"
	"github.com/tietiexx/bot/code/backend/utils/response"
)

func UserLogin(c *gin.Context) {
	var params models.LoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}
	// 查询账号
	user, err := services.UserByUsername(params.Username)
	if err != nil {
		response.UnauthorizedJsonResponse(c, "Invalid credentials")
		return
	}

	err = services.Authenticate(user, params.Password)
	if err != nil {
		response.UnauthorizedJsonResponse(c, "Invalid credentials")
		return
	}

	accessToken, err := jwtutil.GenerateAccessToken(int(user.ID))
	if err != nil {
		response.ErrorJsonResponse(c, "Failed to generate access token")
		return
	}

	refreshToken, err := jwtutil.GenerateRefreshToken(int(user.ID))
	if err != nil {
		response.ErrorJsonResponse(c, "Failed to generate refresh token")
		return
	}

	resp := models.LoginResponse{
		Access:  accessToken,
		Refresh: refreshToken,
	}

	response.SuccessJsonResponse(c, resp)
	return
}

func UserLogout(c *gin.Context) {
	// 退出登录
	response.SuccessJsonResponse(c, nil)
	return
}

func UserInfo(c *gin.Context) {
	userInfo := utils.ContextGetUser(c)
	resp := converters.UserToUserInfoResponse(userInfo)
	response.SuccessJsonResponse(c, resp)
	return
}
