/*
@Time : 2020/11/5 下午11:34
@Author : hoastar
@File : dashboard
@Software: GoLand
*/

package dashboard

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/pkg/pagination"
	"github.com/hoastar/orange/pkg/service"
	"github.com/hoastar/orange/tools/app"
)

func InitData(c *gin.Context) {
	var (
		err        error
		panelGroup struct {
			UserTotalCount      int `json:"user_total_count"`
			WorkOrderTotalCount int `json:"work_order_total_count"`
			UpcomingTotalCount  int `json:"upcoming_total_count"`
			MyUpcomingCount     int `json:"my_upcoming_count"`
		}
		result              interface{}
		processOrderList    []process.Info
		processOrderListMap map[string][]interface{}
	)

	// 查询用户总数
	err = orm.Eloquent.Model(&system.SysUser{}).Count(&panelGroup.UserTotalCount).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 查询工单总数
	err = orm.Eloquent.Model(&process.WorkOrderInfo{}).Count(&panelGroup.WorkOrderTotalCount).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 查询待办总数
	err = orm.Eloquent.Model(&process.WorkOrderInfo{}).
		Where("is_end = 0").
		Count(&panelGroup.UpcomingTotalCount).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 查询我的待办
	w := service.WorkOrder{
		Classify: 1,
		GinObj:   c,
	}
	result, err = w.PureWorkOrderList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	panelGroup.MyUpcomingCount = result.(*pagination.Paginator).TotalCount

	// 查询周工单统计
	statisticsData, err := service.WeeklyStatistics()
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询周工单统计失败，%v", err.Error()))
		return
	}

	// 查询工单提交排名
	submitRankingData, err := service.SubmitRanking()
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询工单提交排名失败，%v", err.Error()))
		return
	}

	// 查询最常用的流程
	err = orm.Eloquent.Model(&process.Info{}).Order("submit_count desc").Limit(10).Find(&processOrderList).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("查询最常用的流程失败，%v", err.Error()))
		return
	}
	processOrderListMap = make(map[string][]interface{})
	for _, v := range processOrderList {
		processOrderListMap["title"] = append(processOrderListMap["title"], v.Name)
		processOrderListMap["submit_count"] = append(processOrderListMap["submit_count"], v.SubmitCount)
	}

	app.Ok(c, map[string]interface{}{
		"panelGroup":        panelGroup,
		"statisticsData":    statisticsData,
		"submitRankingData": submitRankingData,
		"processOrderList":  processOrderListMap,
	}, "")
}