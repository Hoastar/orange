/*
@Time : 2020/11/1 下午10:54
@Author : hoastar
@File : process
@Software: GoLand
*/

package process

import (
	"encoding/json"
	"github.com/hoastar/orange/models/base"
)

// 流程信息
type Info struct {
	base.Model
	// 流程名称
	Name string	`gorm:"column:name; type: varchar(128)" json:"name" form:"name"`
	// 流程图标
	Icon string `gorm:"column:icon; type: varchar(128)" json:"icon" form:"icon"`
	// 流程结构
	Structure json.RawMessage `gorm:"structure; type:json" json:"structure" form:"icon"`
	// 分类ID
	Classify int `gorm:"column:classify; type: int(11)" json:"classify" form:"classify"`
	// 模板
	Tpls json.RawMessage `gorm:"column:tpls; type: json" json:"tpls" form:"tpls"`
	// 任务ID， array
	Task json.RawMessage `gorm:"column:task; type: json" json:"task" form:"task"`
	// 提交统计
	SubmitCount int `gorm:"column:submit_count; type: int(11); default:0" json:"submit_count" form:"submit_count"`
	// 创建者
	Creator int `gorm:"column"reator; type: int(11)" json:"creator" form:"creator"`
	// 绑定通知
	Notice json.RawMessage `gorm:"column:notice; type: json" json:"notice" form:"notice"`
	// 流程备注
	Remarks     string          `gorm:"column:remarks; type:varchar(1024)" json:"remarks" form:"remarks"`
}

func (Info) TableName() string {
	return "p_process_info"
}




























