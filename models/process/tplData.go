/*
@Time : 2020/11/1 下午10:04
@Author : hoastar
@File : tplData
@Software: GoLand
*/

package process

import (
	"encoding/json"
	"github.com/hoastar/orange/models/base"
)

// 工单绑定模板数据
type TplData struct {
	base.Model
	// 工单id
	WorkOrder int `gorm:"column:work_order; type: int(11)" json:"work_order" form:"work_order"`
	// 表单结构
	FormStructure json.RawMessage `gorm:"form_structure; type: json" json:"form_structure" form:"form_structure"`
	// 表单数据
	FormData json.RawMessage `gorm:"column:form_data; type: json" json:"form_data" form:"form_data"`
}

func (TplData) TableName() string {
	return "p_work_order_tpl_data"
}
