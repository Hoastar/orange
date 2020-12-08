/*
@Time : 2020/11/28 下午5:08
@Author : hoastar
@File : workOrder
@Software: GoLand
*/

package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/pkg/service"
	"github.com/hoastar/orange/tools"
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

// 工单列表
func WorkOrderList(c *gin.Context) {
	/*
	1. 代办工单
	2. 我创建的
	3. 我相关的
	4. 所有工单
	 */

	var (
		result 		interface{}
		err			error
		classifyInt int
	)

	classify := c.DefaultQuery("classify", "")
	if classify == "" {
		app.Error(c, -1, errors.New("参数错误，请确认classify是否传递"), "")
		return
	}

	classifyInt, _ = strconv.Atoi(classify)
	w := service.WorkOrder{
		Classify: classifyInt,
		GinObj: c,
	}

	result, err = w.WorkOrderList()
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询工单失败， %v", err.Error()))
		return
	}

	app.Ok(c, result, "")
}

// 处理工单
func ProcessWorkOrder(c *gin.Context) {
	var (
		err				error
		userAuthority	bool
		handle			service.Handle
		params			struct{

			Tasks		[]string
			TargetState	string	`json:"target_state"` // 目标状态
			SourceState string	`json:"source_state"` // 源状态
			WorkOrderId	int		`json:"work_order_id"` // 工单ID
			Circulation string	`json:"circulation"`	// 流转ID
			FlowProperties int	`json:"flow_properties"` // 流转类型 0拒绝，1同意，2其他
			Remarks     string	`json:"remarks"`		// 处理的备注信息
			Tpls        []map[string]interface{}		// 表单数据
		}
	)

	err = c.ShouldBind(&params)
	if err != nil {
		app.Error(c,-1, err, "")
		return
	}

	// 处理工单
	userAuthority, err = service.JudgeUserAuthority(c, params.WorkOrderId, params.SourceState)
	if !userAuthority {
		app.Error(c, -1, errors.New("当前用户没有权限进行此操作"), "")
		return
	}

	err = handle.HandleWorkOrder(
		c,
		params.WorkOrderId,	// 工单ID
		params.Tasks, // 任务列表
		params.TargetState, // 目标节点
		params.SourceState,	// 源节点
		params.Circulation, // 流转标题
		params.FlowProperties, // 流转属性
		params.Remarks, // 备注信息
		params.Tpls, // 工单数据更新
		)

	if err != nil {
		app.Error(c, -1, nil, fmt.Sprintf("处理工单失败， %v", err.Error()))
		return
	}
	app.Ok(c, nil, "工单处理完成")
}

// 结束工单
func UnityWorkOrder(c *gin.Context) {
	var (
		err		error
		workOrderId string
		workOrderInfo process.WorkOrderInfo
		userInfo	system.SysUser
	)

	workOrderId = c.DefaultQuery("work_order_id", "")
	if workOrderId == "" {
		app.Error(c, -1, errors.New("参数不正确， work_order_id"), "")
		return
	}

	tx := orm.Eloquent.Begin()

	// 查询工单信息
	err = tx.Model(&workOrderInfo).
		Where("id = ?", workOrderId).
		Find(&workOrderInfo).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询工单失败， %v", err.Error()))
		return
	}

	if workOrderInfo.IsEnd == 1 {
		app.Error(c, -1, errors.New("工单已经结束"), "")
		return
	}

	// 更新工单状态
	err = tx.Model(&process.WorkOrderInfo{}).
		Where("id = ?", workOrderInfo).
		Update("is_end", 1).Error
	if err != nil {
		tx.Rollback()
		app.Error(c, -1, err, fmt.Sprintf("结束工单失败， %v", err.Error()))
		return
	}

	// 获取当前用户信息
	err = tx.Model(&userInfo).
		Where("user_id = ?", tools.GetUserId(c)).
		Find(&userInfo).Error
	if err != nil {
		tx.Rollback()
		app.Error(c, -1, err, fmt.Sprintf("当前用户查询失败， %v", err.Error()))
		return
	}

	// 写入历史
	tx.Create(&process.CirculationHistory{
		Title: workOrderInfo.Title,
		WorkOrder: workOrderInfo.Id,
		State: "结束工单",
		Circulation: "结束",
		ProcessorId: tools.GetUserId(c),
		Remarks: "手动结束工单。",
	})
	tx.Commit()

	app.Ok(c, nil, "工单已经结束")
}

// 转交工单
func InversionWorkOrder(c *gin.Context) {
	var (
		err				error
		workOrderInfo 	process.WorkOrderInfo
		stateList		[]map[string]interface{}
		stateValue		[]byte
		currentState	map[string]interface{}
		userInfo		system.SysUser
		currentUserInfo	system.SysUser
		params struct{
			WorkOrderId int 	`json:"work_order_id"`
			NodeId		string  `json:"node_id"`
			UserId 		int 	`json:"user_id"`
			Remarks		string	`json:"remarks"`
		}
	)

	// 获取当前用户信息
	err = orm.Eloquent.Model("&currentUserInfo").
		Where("user_id = ?", tools.GetUserId(c)).
		Find(&currentUserInfo).Error

	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询当前用户失败， %v", err.Error()))
		return
	}

	err = c.ShouldBind(&params)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 查询工单信息
	err = orm.Eloquent.Model(&workOrderInfo).
		Where("id = ?", params.WorkOrderId).
		Find(&workOrderInfo).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询工单信息失败， %v", err.Error()))
		return
	}

	// 序列化节点数据
	err = json.Unmarshal(workOrderInfo.State, &stateList)
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("节点数据反序列化失败，%v", err.Error()))
		return
	}

	for _, s := range stateList {
		if s["id"].(string) == params.NodeId {
			s["processor"] = []interface{}{params.UserId}
			s["process_method"] = "person"
			currentState = s
			break
		}
	}

	stateValue, err = json.Marshal(stateList)
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("节点数据序列化失败, %v", err.Error()))
		return
	}
	tx := orm.Eloquent.Begin()

	// 更新数据
	err = tx.Model(&process.WorkOrderInfo{}).
		Where("id = ?", params.WorkOrderId).
		Update("state", stateValue).Error

	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("更新节点信息失败， %v", err.Error()))
		return
	}

	// 查询用户信息
	err = tx.Model(&system.SysUser{}).
		Where("user_id = ?", params.UserId).
		Find(&userInfo).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询用户信息失败，%v", err.Error()))
		return
	}

	// 添加转交历史
	tx.Create(&process.CirculationHistory{
		Title: workOrderInfo.Title,
		WorkOrder: workOrderInfo.Id,
		State: currentState["label"].(string),
		Circulation: "转交",
		Processor: currentUserInfo.NickName,
		ProcessorId: tools.GetUserId(c),
		Remarks: fmt.Sprintf("此阶段负责人已转交给《%v》", userInfo.NickName),
	})
	tx.Commit()
	app.Ok(c, nil, "工单已手动结单")
}

// 催办工单
