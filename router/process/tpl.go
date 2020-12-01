/*
@Time : 2020/11/30 下午10:29
@Author : hoastar
@File : tpl
@Software: GoLand
*/

package process

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/process"
	"github.com/hoastar/orange/middlerware"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterTplRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	tplRouter := v1.Group("/tpl").Use(authMiddleware.MiddlewareFunc()).Use(middlerware.AuthCheckRole())
	{
		tplRouter.GET("", process.TemplateList)
		tplRouter.POST("", process.CreateTemplate)
		tplRouter.PUT("", process.UpdateTemplate)
		tplRouter.DELETE("", process.DeleteTemplate)
		tplRouter.GET("/details", process.TemplateDetails)
	}
}