/*
@Time : 2020/10/29 上午12:16
@Author : hoastar
@File : base
@Software: GoLand
*/

package base

import "github.com/hoastar/orange/pkg/jsonTime"

type Model struct {
	Id int	`gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt jsonTime.JSONTime	`gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdatedAt jsonTime.JSONTime `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt jsonTime.JSONTime	`gorm:"column:delete_time" sql:"index" json:'-'`
}