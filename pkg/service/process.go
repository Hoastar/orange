/*
@Time : 2020/11/8 下午1:43
@Author : hoastar
@File : process
@Software: GoLand
*/

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/tools"
	"strconv"
)

type WorkOrderData struct {
	// 工单基本信息
	process.WorkOrderInfo
	// 工单当前状态
	CurrentState string `json:"current_state"`
}


func ProcessStructure(c *gin.Context, processId int, workOrderId int) (result map[string]interface{}, err error) {
	var (
		// 流程基本信息
		processValue			process.Info
		// 流程结构化详情
		processStructureDetails map[string]interface{}
		// 工单节点
		processNode				[]map[string]interface{}
		// 模板详情
		tplDetails				[]*process.TplInfo
		// 工单信息
		workOrderInfo			WorkOrderData
		// 工单（所选择的流程）模板集合
		workOrderTpls			[]*process.TplData
		// process.CirculationHistory 工单流转历史
		workOrderHistory		[]*process.CirculationHistory
		// 工单节点状态集合
		stateList				[]map[string]interface{}
	)

	err = orm.Eloquent.Model(&processValue).Where("id = ?", processId).Find(&processValue).Error
	if err != nil {
		err = fmt.Errorf("查询流程失败， %v", err.Error())
		return
	}

	err = json.Unmarshal([]byte(processValue.Structure), &processStructureDetails)
	if err != nil {
		err = fmt.Errorf("json转map失败， %v", err.Error())
		return
	}

	// 排序，冒泡排序
	p := processStructureDetails["nodes"].([]interface{})
	if len(p) > 1 {
		for i := 0; i < len(p); i++ {
			for j := 1; j < len(p)-i; j++ {
				if p[j].(map[string]interface{})["sort"] ==  nil || p[j-1].(map[string]interface{})["sort"] == nil {
					return nil, errors.New("流程未定义属性， 请确认")
				}

				leftInt, _ := strconv.Atoi(p[j].(map[string]interface{})["sort"].(string))
				rightInt, _ := strconv.Atoi(p[j-1].(map[string]interface{})["sort"].(string))

				if leftInt < rightInt {
					// 交换
					p[j], p[j-1] = p[j-1], p[j]
				}
			}

		}

		for _, node := range processStructureDetails["nodes"].([]interface{}) {
			processNode = append(processNode, node.(map[string]interface{}))
		}
	} else {
		processNode = processStructureDetails["nodes"].([]map[string]interface{})
	}

	processValue.Structure = nil
	result = map[string]interface{}{
		// 流程基本信息
		"process": processValue,
		// 流程节点
		"nodes": processNode,
		// 流程结构中的流程线
		"edges": processStructureDetails["edges"],
	}


	// 获取历史记录
	err = orm.Eloquent.Model(&process.CirculationHistory{}).
		Where("work_order = ?", workOrderId).
		Order("id desc").
		Find(&workOrderHistory).Error

	if err != nil {
		return
	}

	result["circulationHistory"] = workOrderHistory

	if workOrderId == 0 {
		// 查询流程模板
		var tplIdList []int
		err = json.Unmarshal(processValue.Tpls, &tplIdList)
		if err != nil {
			err = fmt.Errorf("json转map失败， %v", err.Error())
			return
		}

		err = orm.Eloquent.Model(&tplDetails).
			Where("id in (?)", tplIdList).
			Find(&tplDetails).Error
		if err != nil {
			err = fmt.Errorf("查询模板失败，%v", err.Error())
			return
		}
		result["tpls"] = tplDetails
	} else {
		// 查询工单信息
		err = orm.Eloquent.Model(&process.WorkOrderInfo{}).
			Where("id = ?", workOrderId).
			Scan(&workOrderInfo).Error
		if err != nil {
			return
		}

		// 获取当前节点

		err = json.Unmarshal(workOrderInfo.State, &stateList)
		if err != nil {
			err = fmt.Errorf("序列化节点列表失败， %v", err.Error())
			return
		}

		if len(stateList) == 0 {
			err = errors.New("当前工单没有下一节点数据")
			return
		}

		// 整理需要并行处理的数据
		if len(stateList) > 1 {
		continueHistoryTag:
			for _, v := range workOrderHistory {
				status := false
				for i, s := range stateList {
					if v.Source == s["id"].(string) && v.Target != "" {
						status = true
						stateList = append(stateList[:i], stateList[i+1:]...)
						continue continueHistoryTag
					}
				}
				if !status {
					break
				}
			}
		}

		if len(stateList) > 0 {
		breakStateTag:
			for _, stateValue := range stateList {
				for _, processNodeValue := range processStructureDetails["nodes"].([]interface{}) {
					if stateValue["id"].(string) == processNodeValue.(map[string]interface{})["id"] {
						if _, ok := stateValue["processor"]; ok {
							for _, userId := range stateValue["processor"].([]interface{}) {
								if int(userId.(float64)) == tools.GetUserId(c) {
									workOrderInfo.CurrentState = stateValue["id"].(string)
									break breakStateTag
								}
							}
						} else {
							err = errors.New("未查询到对应的处理人字段，请确认。")
							return
						}
					}
				}
			}
			if workOrderInfo.CurrentState == "" {
				workOrderInfo.CurrentState = stateList[0]["id"].(string)
			}
		}
		result["workOrder"] = workOrderInfo

		// 查询工单表单数据
		err = orm.Eloquent.Model(&workOrderTpls).
			Where("work_order = ?", workOrderId).
			Find(&workOrderTpls).Error

		if err != nil {
			return nil, err
		}

		result["tpls"] = workOrderTpls
	}
	return result, nil
}
