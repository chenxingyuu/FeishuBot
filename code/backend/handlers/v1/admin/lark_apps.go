package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tietiexx/bot/code/backend/converters"
	"github.com/tietiexx/bot/code/backend/models"
	"github.com/tietiexx/bot/code/backend/services"
	"github.com/tietiexx/bot/code/backend/utils/response"
	"strings"
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

	uuidStr := strings.ReplaceAll(uuid.NewString(), "-", "")
	err := services.LarkAppCreate(uuidStr, params.Name)
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
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppUpdatePartial(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}

func LarkAppDelete(c *gin.Context) {
	response.SuccessJsonResponse(c, nil)
	return
}
