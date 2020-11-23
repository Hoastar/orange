/*
@Time : 2020/11/1 下午8:36
@Author : hoastar
@File : login
@Software: GoLand
*/

package system

import (
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/tools"
)

// Login 登录信息实体
type Login struct {
	Username  string `form:"UserName" json:"username" binding:"required"`
	Password  string `form:"Password" json:"password" binding:"required"`
	Code      string `form:"Code" json:"code" binding:"required"`
	UUID      string `form:"UUID" json:"uuid" binding:"required"`
	LoginType int    `form:"LoginType" json:"loginType"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, err error) {
	err = orm.Eloquent.Table("sys_user").Where("username = ? ", u.Username).Find(&user).Error
	if err != nil {
		return
	}

	// 验证密码正确性
	if u.LoginType == 0 {
		_, err = tools.CompareHashAndPassword(user.Password, u.Password)
		if err != nil {
			return
		}
	}

	err = orm.Eloquent.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		return
	}
	return
}



























