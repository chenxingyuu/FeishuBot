package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/tietiexx/bot/code/backend/converters"
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/models"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils/response"
)

func LarkAppList(c *gin.Context) {
	var params models.LarkAppListRequest
	err := c.ShouldBindQuery(&params)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	count, list := services.LarkAppListPaginate(params)

	resp := converters.ToLarkAppPaginationResponse(count, list)
	response.SuccessJsonResponse(c, resp)
	return
}

func LarkAppCreate(c *gin.Context) {
	var params models.LarkAppCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	err := services.LarkAppCreate(params.Name)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppDetail(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppUpdate(c *gin.Context) {
	larkAppUUID := c.Param("larkAppUUID")

	var params models.LarkAppUpdateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	larkApp, err := services.LarkAppByUUID(larkAppUUID)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	larkApp.Name = params.Name
	larkApp.Status = params.Status
	larkApp.AppId = params.AppID
	larkApp.AppSecret = params.AppSecret
	larkApp.EncryptKey = params.EncryptKey
	larkApp.VerificationToken = params.VerificationToken

	err = services.LarkAppSave(larkApp)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppUpdatePartial(c *gin.Context) {
	larkAppUUID := c.Param("larkAppUUID")

	var params models.LarkAppUpdatePartialRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	_, err := services.LarkAppByUUID(larkAppUUID)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	larkApp := database.LarkApp{Name: params.Name, Status: params.Status, AppId: params.AppID, AppSecret: params.AppSecret, EncryptKey: params.EncryptKey, VerificationToken: params.VerificationToken}
	err = services.LarkAppUpdateByUUID(larkAppUUID, &larkApp)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}

	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppDelete(c *gin.Context) {
	larkAppUUID := c.Param("larkAppUUID")
	err := services.LarkAppDelete(larkAppUUID)
	if err != nil {
		response.ErrorJsonResponse(c, err.Error())
		return
	}
	response.SuccessJsonResponse(c, nil)
	return
}
