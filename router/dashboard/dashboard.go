/*
@Time : 2020/11/28 下午7:08
@Author : hoastar
@File : dashboard
@Software: GoLand
*/

package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/dashboard"
	"github.com/hoastar/orange/middleware"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	classify := v1.Group("/dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		classify.GET("", dashboard.InitData)
	}
}
