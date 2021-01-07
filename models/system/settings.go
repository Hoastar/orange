/*
@Time : 2020/11/1 下午8:41
@Author : hoastar
@File : settings
@Software: GoLand
*/

package system

import (
	"encoding/json"
	"github.com/hoastar/orange/models/base"
)

type Settings struct {
	base.Model
	// 设置分类，1 配置信息，2 Ldap配置
	Classify int             `gorm:"column:classify; type:int(11)" json:"classify" form:"classify"`
	// 配置内容
	Content  json.RawMessage `gorm:"column:content; type:json" json:"content" form:"content"`
}

func (Settings) TableName() string {
	return "sys_settings"
}
