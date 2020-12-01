/*
@Time : 2020/11/28 下午10:51
@Author : hoastar
@File : process
@Software: GoLand
*/

package process

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/process"
	"github.com/hoastar/orange/middlerware"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterProcessRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	processRouter := v1.Group("/process").Use(authMiddleware.MiddlewareFunc()).Use(middlerware.AuthCheckRole())
	{
		processRouter.GET("/classify", process.ClassifyProcessList)
		processRouter.GET("", process.ProcessList)
		processRouter.POST("", process.CreateProcess)
		processRouter.PUT("", process.UpdateProcess)
		processRouter.DELETE("", process.DeleteProcess)
		processRouter.GET("/details", process.ProcessDetails)
	}
}