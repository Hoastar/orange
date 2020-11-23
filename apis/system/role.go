/*
@Time : 2020/11/16 上午10:39
@Author : hoastar
@File : role
@Software: GoLand
*/

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/tools"
	"github.com/hoastar/orange/tools/app"
)

// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色/Role
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/rolelist [get]
// @Security
func GetRoleList(c *gin.Context) {
	var (
		err 		error
		pageSize = 	10
		pageIndex =	1
		data		system.SysRole
	)

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.RoleKey = c.Request.FormValue("roleKey")
	data.RoleName = c.Request.FormValue("roleName")
	data.Status = c.Request.FormValue("status")
	result, count, err :=data.GetPage(pageSize, pageIndex)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param roleId path string false "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/role [get]
// @Security Bearer

func GetRole(c *gin.Context) {
	var (
		err		error
		Role	system.SysRole
	)
	Role.RoleId, _ = tools.StringToInt(c.Param("roleId"))
	result, err := Role.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	menuIds, err := Role.GetRoleMenu()

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	result.MenuIds = menuIds
	app.Ok(c, result, "")

}

// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body models.SysRole true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/role [post]

func InsertRole(c *gin.Context) {
	var data system.SysRole
	data.CreateBy = tools.GetUserIdToStr(c)
	err := c.BindWith(&data, binding.JSON)

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	id, err := data.Insert()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	data.RoleId = id
	var t system.RoleMenu
	_, err = t.Insert(id, data.MenuIds)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.Ok(c, data, "添加成功")
}


























