/*
@Time : 2020/11/3 下午9:58
@Author : hoastar
@File : logger
@Software: GoLand
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/pkg/logger"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.Method

		// 状态码
		statusCode := c.Writer.Status()

		// 请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof(" %s %3d %13v %15s %s %s",
			startTime.Format("2006-01-02 15:04:05.9999"),
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			)
	}
}
