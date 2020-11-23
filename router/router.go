/*
@Time : 2020/11/3 上午12:41
@Author : hoastar
@File : router
@Software: GoLand
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/handler"
	config2 "github.com/hoastar/orange/tools/config"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	if config2.ApplicationConfig.IsHttps {
		r.Use(handler.TlsHandler())
	}
	middleware.InitMiddleware(r)
}
