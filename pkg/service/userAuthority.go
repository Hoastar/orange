/*
@Time : 2020/11/8 下午5:20
@Author : hoastar
@File : userAuthority
@Software: GoLand
*/

package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/tools"
)

func JudgeUserAuthority(c *gin.Context, workOrderId int, currentState string) (status bool, err error) {
	/*
		person 人员
		persongroup 人员组
		department 部门
		variable 变量
	*/

	var (
		userDept			system.Dept
		workOrderInfo		process.WorkOrderInfo
		userinfo			system.SysUser
		cirHistoryList		[]process.CirculationHistory
		stateValue			map[string]interface{}
		processInfo			process.Info
		processState		ProcessState
		currentStateList	[]map[string]interface{}
		currentStateValue	map[string]interface{}
	)

	// 获取工单信息
	err = orm.Eloquent.Model(&workOrderInfo).
		Where("id = ?", workOrderId).
		Find(&workOrderInfo).Error

	if err != nil {
		return
	}

	// 获取流程信息
	err = orm.Eloquent.Model(&process.Info{}).
		Where("id = ?", workOrderInfo.Process).Find(&processInfo).Error

	if err != nil {
		return
	}

	err = json.Unmarshal(processInfo.Structure, &processState.Structure)
	if err != nil {
		return
	}

	stateValue, err = processState.GetNode(currentState)

	if err != nil {
		return
	}

	err = json.Unmarshal(workOrderInfo.State, &currentStateList)
	if err != nil {
		return
	}

	for _, v := range currentStateList {
		if v["id"].(string) == currentState {
			currentStateValue = v
			break
		}
	}

	// 会签

	if currentStateValue["processor"] != nil && len(currentStateValue["processor"].([]interface{})) > 1 {
		if isCounterSign, ok := stateValue["isCounterSign"]; ok {
			if isCounterSign.(bool) {
				err = orm.Eloquent.Model(&process.CirculationHistory{}).
					Where("work_order = ?", workOrderId).
					Find(&cirHistoryList).Error
				if err != nil {
					return
				}

				for _, cirHistoryValue := range cirHistoryList {
					if cirHistoryValue.Source != stateValue["id"] {
						break
					}

					if cirHistoryValue.Source == stateValue["id"] && cirHistoryValue.ProcessorId == tools.GetUserId(c) {
						return
					}
				}
			}
		}
	}

	switch currentStateValue["process_method"].(string) {
	case "person":
		for _, processorValue := range currentStateValue["processor"].([]interface{}) {
			if int(processorValue.(float64)) == tools.GetUserId(c) {
				status = true
			}
		}


	case "variable":
		for _, p := range currentStateValue["processor"].([]interface{}) {
			switch int(p.(float64)) {
			case 1:
				if workOrderInfo.Creator == tools.GetUserId(c) {
					status = true
				}
			case 2:
				err = orm.Eloquent.Model(&userinfo).Where("user_id = ?", workOrderInfo.Creator).Find(&userinfo).Error
				if err != nil {
					return
				}
				err = orm.Eloquent.Model(&userDept).Where("dept_id = ?", userinfo.DeptId).Find(&userDept).Error
				if err != nil {
					return
				}

				if userDept.Leader == tools.GetUserId(c) {
					status = true
				}
			}
		}
	}
	return
}
