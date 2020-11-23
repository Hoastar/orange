/*
@Time : 2020/11/16 上午12:30
@Author : hoastar
@File : rolemenu
@Software: GoLand
*/

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/tools/app"
	"net/http"
)

// @Summary RoleMenu列表数据
// @Description 获取JSON
// @Tags 角色菜单
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/rolemenu [get]
// @Security Bearer

func GetRoleMenu(c *gin.Context) {
	var (
		res app.Response
		Rm 	system.RoleMenu
	)
	_ = c.ShouldBind(&Rm)
	result, err := Rm.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}