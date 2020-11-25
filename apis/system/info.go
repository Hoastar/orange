/*
@Time : 2020/11/23 上午11:16
@Author : hoastar
@File : info
@Software: GoLand
*/

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/tools"
	"github.com/hoastar/orange/tools/app"
)

func GetInfo(c *gin.Context)  {
	var roles = make([]string, 1)
	roles[0] = tools.GetRoleName(c)

	var permissons = make([]string, 1)
	permissons[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	RoleMenu := system.RoleMenu{}
	RoleMenu.RoleId = tools.GetRoleId(c)

	var mp = make(map[string]interface{})

	mp["roles"] = roles
	if tools.GetRoleName(c) == "admin" || tools.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissons
		mp["buttons"] = buttons
	} else {
		list, _ := RoleMenu.GetPermis()
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysuser := system.SysUser{}
	sysuser.UserId = tools.GetUserId(c)

	user, err := sysuser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	mp["avatar"] = "https://https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["userName"] = user.NickName
	mp["userId"] = user.UserId
	mp["deptId"] = user.DeptId
	mp["name"] = user.NickName

	app.Ok(c, mp, "")
}
