/*
@Time : 2020/11/13 下午9:53
@Author : hoastar
@File : loginLog
@Software: GoLand
*/

package log

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/system"
	"github.com/hoastar/orange/tools"
	"github.com/hoastar/orange/tools/app"
	"net/http"
)

// @Summary 登录日志列表
// @Description 获取JSON
// @Router /api/v1/loginloglist [get]

func GetLoginLogList(c *gin.Context) {
	var (
		err				error
		pageSize 	= 	10
		pageIndex	= 	1
		data			system.LoginLog
	)

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	data.Username = c.Request.FormValue("username")
	data.Status = c.Request.FormValue("status")
	data.Ipaddr = c.Request.FormValue("ipaddr")
	result, count, err := data.GetPage(pageSize, pageIndex)

	if err != nil {
		app.Error(c, -1, err, "")
	}

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 添加日志
// @Description 获取JSON
// @Router /api/v1/loginloglist [post]

func InsertLoginLog(c *gin.Context) {
	var data system.LoginLog
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	result, err := data.Create()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改登录日志
// @Description 获取JSON
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/loginlog [put]

func UpdateLoginLog(c *gin.Context) {
	var (
		res app.Response
		data system.LoginLog
	)
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	result, err := data.Update(data.InfoId)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}


// @Summary 批量删除登录日志
// @Description 删除数据
// @Param infoId path string true "以逗号（,）分割的infoId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/loginlog/{infoId} [delete]
func DeleteLoginLog(c *gin.Context)  {
	var (
		res app.Response
		data system.LoginLog
	)

	data.UpdateBy = tools.GetUserIdToStr(c)
	IDS := tools.IdsStrToIdsIntGroup("infoId", c)

	_, err := data.BatchDelete(IDS)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func CleanLoginLog(c *gin.Context) {
	err := orm.Eloquent.Delete(&system.LoginLog{}).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, "", "已清空")
}
