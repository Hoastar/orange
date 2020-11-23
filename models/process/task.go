/*
@Time : 2020/11/1 下午11:44
@Author : hoastar
@File : task
@Software: GoLand
*/

package process

import "github.com/hoastar/orange/models/base"

// 任务
type TaskInfo struct {
	base.Model
	// TaskName
	Name string `gorm:"column:name; type: varchar(256)" json:"name" form:"name"`
	// TaskType
	Type string `gorm:"column:task_type; type: varchar(45)" json:"task_type" form:"task_type"`
	// TaskConent
	Content  string `gorm:"column:content; type: longtext" json:"content" form:"content"`
	// 创建者
	Creator  int    `gorm:"column:creator; type: int(11)" json:"creator" form:"creator"`
	// 备注
	Remarks string `gorm:"column:remarks; type: longtext" json:"remarks" form:"remarks"`
}

func (TaskInfo) TableName() string {
	return "p_task_info"
}
