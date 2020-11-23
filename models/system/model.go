/*
@Time : 2020/11/1 下午8:20
@Author : hoastar
@File : model
@Software: GoLand
*/

package system

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:creat_time" json:"creat_time" form:"create_time"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt *time.Time `gorm:"column:delete_time" sql:"index" json:"-"`
}
