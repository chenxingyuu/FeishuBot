package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type JsonResponseInfo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	LogId   string      `json:"log_id"`
}

func FileResponse(c *gin.Context, filePath string) {
	fileName := filepath.Base(filePath)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}

func JsonResponse(c *gin.Context, code int, data JsonResponseInfo) {
	data.LogId = c.GetString("logId")
	c.JSON(code, data)
}

func SuccessJsonResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, JsonResponseInfo{Code: 0, Data: data})
}

func ErrorJsonResponse(c *gin.Context, message string) {
	JsonResponse(c, http.StatusBadRequest, JsonResponseInfo{Code: -1, Message: message})
}

func UnauthorizedJsonResponse(c *gin.Context, message string) {
	JsonResponse(c, http.StatusUnauthorized, JsonResponseInfo{Code: -1, Message: message})
}
