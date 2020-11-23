/*
@Time : 2020/11/3 上午12:53
@Author : hoastar
@File : init
@Software: GoLand
*/

package middlerware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// Set X-Request-Id header
	r.Use(RequestId())
}

