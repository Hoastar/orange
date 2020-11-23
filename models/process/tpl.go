/*
@Time : 2020/11/1 下午9:19
@Author : hoastar
@File : tpl
@Software: GoLand
*/

package process

import (
	"encoding/json"
	"github.com/hoastar/orange/models/base"
)

type TplInfo struct {
	base.Model
	// 模板名称
	Name	string `gorm:"column:name; type: varchar(128)" json:"name" form:"name" binding:"required"`
	// 表单结构
	FormStructure json.RawMessage `gorm:"column:form_structure; type: json" json:"form_structure" form:"form_structure" binding:"required"`
	Creator int `gorm:"column:creator; type: int(11)" json:"creator" form:"creator"`
	Remarks string `gorm:"column:remarks; type: longtext" json:"remarks" form:"remarks"`
}

func (TplInfo) TableName() string {
	return "p_tpl_info"
}
