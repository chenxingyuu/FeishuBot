package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/converters"
	"github.com/tietiexx/bot/code/backend/models"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils/response"
)

func LarkBotList(c *gin.Context) {
	var params models.LarkBotListRequest
	err := c.ShouldBindQuery(&params)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	count, list := services.LarkBotListPaginate(params)

	resp := converters.ToLarkBotPaginationResponse(count, list)
	response.SuccessJsonResponse(c, resp)
	return
}

func LarkBotCreate(c *gin.Context) {
	var params models.LarkBotCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	larkApp, err := services.LarkAppByUUID(params.LarkAppUUID)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	err = services.LarkBotCreate(larkApp.ID, params.Name, params.LarkBotType)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	response.SuccessJsonResponse(c, nil)
	return
}

func LarkBotDetail(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkBotUpdate(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkBotUpdatePartial(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkBotDelete(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}
