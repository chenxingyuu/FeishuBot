package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	greenB   = "\033[97;42m"
	whiteB   = "\033[90;47m"
	yellowB  = "\033[97;43m"
	redB     = "\033[97;41m"
	blueB    = "\033[97;44m"
	magentaB = "\033[97;45m"
	cyanB    = "\033[97;46m"
)
const (
	greenF   = "\033[32m"
	whiteF   = "\033[37m"
	yellowF  = "\033[33m"
	redF     = "\033[31m"
	blueF    = "\033[34m"
	magentaF = "\033[35m"
	cyanF    = "\033[36m"
)

const reset = "\033[0m"

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return greenB
	case code >= 300 && code < 400:
		return cyanB
	case code >= 400 && code < 500:
		return yellowB
	default:
		return redB
	}
}
func colorForMethod(method string) string {
	switch {
	case method == http.MethodGet:
		return greenF
	case method == http.MethodPost:
		return yellowF
	case method == http.MethodPut || method == http.MethodPatch:
		return blueF
	case method == http.MethodDelete:
		return redF
	case method == http.MethodOptions || method == http.MethodHead:
		return blueF
	default:
		return magentaF
	}
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logId := strconv.FormatInt(time.Now().UnixNano()+int64(rand.Intn(1000)), 10)
		c.Set("logId", logId)
		c.Next()
	}
}

func GinCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logId := c.GetString("logId")
		start := time.Now()
		c.Next()
		latency := time.Now().Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		methodColor := colorForMethod(method)
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)

		fmt.Printf("%s [%4s] [%12v] [%19s] %-16v |%s  %3d  %s|%s %-6s %s| %-7s %s\n",
			start.Format("2006/01/02 15:04:05"),
			"GIN",
			latency,
			logId,
			clientIP,
			statusColor,
			statusCode,
			reset,
			methodColor,
			method,
			reset,
			path,
			c.Errors.String(),
		)
	}
}
