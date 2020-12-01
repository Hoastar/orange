/*
@Time : 2020/11/28 下午5:08
@Author : hoastar
@File : workOrder
@Software: GoLand
*/

package process

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/pkg/service"
	"github.com/hoastar/orange/tools/app"
	"strconv"
)

// 流程结构包括节点、转流和模板
func ProcessStructure(c *gin.Context) {
	processId := c.DefaultQuery("processId", "")
	if processId == "" {
		app.Error(c, -1, errors.New("参数不正确，请确定参数processId是否传递"), "")
		return
	}

	workOrderId := c.DefaultQuery("workOrderId", "0")
	if workOrderId == "" {
		app.Error(c, -1, errors.New("参数不正确，请确定参数workOrderId是否传递"), "")
		return
	}

	workOrderIdInt, _ := strconv.Atoi(workOrderId)
	processIdInt, _ := strconv.Atoi(processId)

	result, err := service.ProcessStructure(c, processIdInt, workOrderIdInt)

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	if workOrderIdInt != 0 {
		currentState := result["workOrder"].(service.WorkOrderData).CurrentState
		userAuthority, err := service.JudgeUserAuthority(c, workOrderIdInt, currentState)
		if err != nil {
			app.Error(c, -1, err, fmt.Sprintf("判断用户是否有权限，失败， %v", err.Error()))
			return
		}
		result["userAuthority"] = userAuthority
	}

	app.Ok(c, result, "数据获取成功")
}

// 新建工单
func CreateWorkOrder(c *gin.Context) {
	err := service.CreateWorkOrder(c)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, "", "成功提交工单时间")
}

