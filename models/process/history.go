/*
@Time : 2020/11/1 下午10:17
@Author : hoastar
@File : history
@Software: GoLand
*/

package process

import (
	"github.com/hoastar/orange/models/base"
)

type History struct {
	base.Model
	// 任务id
	Task int `gorm:"column:task; type: int(11)" json:"task" form:"task"`
	// 任务名称
	Name string `gorm:"column:task; type: varchar(256)" json:"name" form:"name"`
	// 任务类型： python, shell
	TaskType int `gorm:"column:task_type; type: int(11)" json:"task_type" form:"task_type"`
	// 执行时间
	ExecutionTime string `gorm:"column:excution_time; type: varchar(128)" json:"execution_time" form:"execution_time"`
	// 任务结果
	Result string `gorm:"column:result; type: longtext" json:"result" form:"result"`
}

func (History) TableName() string {
	return "p_task_history"
}