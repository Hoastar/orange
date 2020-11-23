/*
@Time : 2020/11/7 下午5:30
@Author : hoastar
@File : workOrderList
@Software: GoLand
*/

package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/pkg/pagination"
	"github.com/hoastar/orange/tools"
)

type WorkOrder struct {
	// Classify
	// 1: 代办工单
	// 2: 我创建的
	// 3: 我相关的
	// 4: 所有工单
	Classify int
	GinObj *gin.Context
}

type workOrderInfo struct {
	process.WorkOrderInfo
	Principals	string `json:"principals"`
	StateName string `json:"state_name"`
	DataClassify int `json:"data_classify"`
}

func (w *WorkOrder) PureWorkOrderList() (result interface{}, err error) {
	var workOrderInfoList []workOrderInfo
	title := w.GinObj.DefaultQuery("title", "")
	db := orm.Eloquent.Model(&process.WorkOrderInfo{}).Where("title like ?", fmt.Sprintf("%%%v%%", title))

	// 获取当前用户信息
	switch w.Classify {
	case 1:
		// 代办工单
		// 1: 个人
		personSelect := fmt.Sprintf("(JSON_CONTAINS(state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(state, JSON_OBJECT('process_method', 'person')))", tools.GetUserId(w.GinObj))
		db = db.Where(fmt.Sprintf("(%v) and is_end = 0", personSelect))
	case 2:
		// 2: 我创建的
		db = db.Where("creator = ?", tools.GetUserId(w.GinObj))
	case 3:
		// 3: 与我相关的
		db = db.Where(fmt.Sprintf("JSON_CONTAINS(related_person, '%v')", tools.GetUserId(w.GinObj)))
	case 4:
		// 所有工单
		default:
		return nil, fmt.Errorf("请确认查询的数据类型是否正确")
	}

	result, err = pagination.Paging(&pagination.Param{
		C: w.GinObj,
		DB: db,
	}, &workOrderInfoList)

	if err != nil {
		err = fmt.Errorf("查询工单列表失败， %v", err.Error())
		return
	}
	return
}

func (w *WorkOrder) WorkOrderList() (result interface{}, err error) {
	var (
		// 当前经手(处理)人
		principals	string
		// result中所有的工单状态集合
		StateList	[]map[string]interface{}
		// 完整的工单信息集合
		workOrderInfoList []workOrderInfo
		minusTotal	int
	)

	result, err = w.PureWorkOrderList()
	if err != nil {
		return
	}

	for i, v := range *result.(*pagination.Paginator).Data.(*[]workOrderInfo) {
		err = json.Unmarshal(v.State, &StateList)
		if err != nil {
			err = fmt.Errorf("json反序列化失败，%v", err.Error())
			return
		}
		var (
			stateName    string
			structResult map[string]interface{}
			authStatus   bool
		)
		if len(StateList) != 0 {
			// 仅待办工单需要验证
			// todo：还需要找最优解决方案
			if w.Classify == 1 {
				structResult, err = ProcessStructure(w.GinObj, v.Process, v.Id)
				if err != nil {
					return
				}

				authStatus, err = JudgeUserAuthority(w.GinObj, v.Id, structResult["workOrder"].(WorkOrderData).CurrentState)
				if err != nil {
					return
				}
				if !authStatus {
					minusTotal += 1
					continue
				}
			} else {
				authStatus = true
			}

			processorList := make([]int, 0)
			if len(StateList) > 1 {
				for _, s := range StateList {
					for _, p := range s["processor"].([]interface{}) {
						if int(p.(float64)) == tools.GetUserId(w.GinObj) {
							processorList = append(processorList, int(p.(float64)))
						}
					}
					if len(processorList) > 0 {
						stateName = s["label"].(string)
						break
					}
				}
			}
			if len(processorList) == 0 {
				for _, v := range StateList[0]["processor"].([]interface{}) {
					processorList = append(processorList, int(v.(float64)))
				}
				stateName = StateList[0]["label"].(string)
			}
			principals, err = GetPrincipal(processorList, StateList[0]["process_method"].(string))
			if err != nil {
				err = fmt.Errorf("查询处理人名称失败，%v", err.Error())
				return
			}
		}
		workOrderDetails := *result.(*pagination.Paginator).Data.(*[]workOrderInfo)
		workOrderDetails[i].Principals = principals
		workOrderDetails[i].StateName = stateName
		workOrderDetails[i].DataClassify = v.Classify
		if authStatus {
			workOrderInfoList = append(workOrderInfoList, workOrderDetails[i])
		}
	}

	result.(*pagination.Paginator).Data = &workOrderInfoList
	result.(*pagination.Paginator).TotalCount -= minusTotal

	return result, nil
}