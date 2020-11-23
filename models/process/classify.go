/*
@Time : 2020/11/1 下午9:13
@Author : hoastar
@File : classify
@Software: GoLand
*/

package process

import "github.com/hoastar/orange/models/base"

// Classify 流程分类
type Classify struct {
	base.Model
	Name string `gorm:"column:name; type: varchar(128)" json:"name" form:"name"` // 分类名称
	Creator int `gorm:"column:creator; type: int(11) json:"creator" form:"creator"` // 创建者
}

func (Classify) TableName() string  {
	return "p_process_classify"
}