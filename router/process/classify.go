/*
@Time : 2020/11/28 下午10:50
@Author : hoastar
@File : classify
@Software: GoLand
*/

package process

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/process"
	"github.com/hoastar/orange/middleware"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterClassifyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	classify := v1.Group("/classify").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		classify.GET("", process.ClassifyList)
		classify.POST("", process.CreateClassify)
		classify.PUT("", process.UpdateClassify)
		classify.DELETE("", process.DeleteClassify)
	}
}