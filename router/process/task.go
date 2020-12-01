/*
@Time : 2020/11/30 下午10:30
@Author : hoastar
@File : task
@Software: GoLand
*/

package process

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/process"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterTaskRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	taskRouter := v1.Group("/task").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		taskRouter.GET("", process.TaskList)
		taskRouter.GET("/details", process.TaskDetails)
		taskRouter.POST("", process.CreateTask)
		taskRouter.PUT("", process.UpdateTask)
		taskRouter.DELETE("", process.DeleteTask)
	}
}