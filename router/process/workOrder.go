/*
@Time : 2020/11/30 下午10:29
@Author : hoastar
@File : workOrder
@Software: GoLand
*/

package process

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/apis/process"
	"github.com/hoastar/orange/middleware"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
)

func RegisterWorkOrderRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	workOrderRouter := v1.Group("/work-order").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		workOrderRouter.GET("/process-structure", process.ProcessStructure)
		workOrderRouter.POST("/create", process.CreateWorkOrder)
		workOrderRouter.GET("/list", process.WorkOrderList)
		workOrderRouter.POST("/handle", process.ProcessWorkOrder)
		workOrderRouter.GET("/unity", process.UnityWorkOrder)
		workOrderRouter.POST("/inversion", process.InversionWorkOrder)
		workOrderRouter.GET("/urge", process.UrgeWorkOrder)
		workOrderRouter.PUT("/active-order/:id", process.ActiveOrder)
		workOrderRouter.DELETE("/delete/:id", process.DeleteWorkOrder)
	}
}