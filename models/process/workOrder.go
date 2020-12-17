/*
@Time : 2020/11/1 下午11:22
@Author : hoastar
@File : workOrder
@Software: GoLand
*/

package process

import (
	"encoding/json"
	"github.com/hoastar/orange/models/base"
)

// WorkOrderInfo 工单
type WorkOrderInfo struct {
	base.Model
	// 工单标题
	Title string `gorm:"column:title; type: varchar(11)" json:"title" form:"title"`
	// 工单优先级 1，正常 2，紧急 3，非常紧急
	Priority int `gorm:"column:priority; type:int(11)" json:"priority" form:"priority"`
	// 工单优先级 1，正常 2，紧急 3，非常紧急
	// 流程ID
	Process int `gorm:"column:process; type:int(11)" json:"process" form:"process"`
	// 分类ID
	Classify int `gorm:"column:classify; type:int(11)" json:"classify" form:"classify"`
	// 是否结束：0 未结束，1 已结束
	IsEnd int `gorm:"column:is_end; type:int(11); default:0" json:"is_end" form:"is_end"`
	// 是否被拒绝： 0 没有，1 有
	IsDenied int `gorm:"column:is_denied; type:int(11); default:0" json:"is_denied" form:"is_denied"`
	// 状态信息
	State json.RawMessage `gorm:"column:state; type:json" json:"state" form:"state"`
	// 工单所有（经手的）处理人
	Creator int `gorm:"column:creator; type:int(11)" json:"creator" form:"creator"`
	// 催办次数
	UrgeCount int `gorm:"column:urge_count; type:int(11); default:0" json:"urge_count" form:"urge_count"`
	// 上一次催促时间
	UrgeLastTime  int `gorm:"column:urge_last_time; type:int(11); default:0" json:"urge_last_time" form:"urge_last_time"`
	RelatedPerson json.RawMessage `gorm:"column:related_person; type:json" json:"related_person" form:"related_person"`
}

func (WorkOrderInfo) TableName() string {
	return "p_work_order_info"
}
