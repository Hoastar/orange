/*
@Time : 2020/11/3 上午12:41
@Author : hoastar
@File : init_router
@Software: GoLand
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/handler"
	"github.com/hoastar/orange/middleware"
	_ "github.com/hoastar/orange/middleware"
	_ "github.com/hoastar/orange/pkg/jwtauth"
	"github.com/hoastar/orange/tools"
	config2 "github.com/hoastar/orange/tools/config"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	if config2.ApplicationConfig.IsHttps {
		r.Use(handler.TlsHandler())
	}
	middleware.InitMiddleware(r)
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	return r
}