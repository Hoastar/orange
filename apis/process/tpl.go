/*
@Time : 2020/11/27 下午3:47
@Author : hoastar
@File : tpl
@Software: GoLand
*/

package process

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/process"
	"github.com/hoastar/orange/pkg/pagination"
	"github.com/hoastar/orange/tools"
	"github.com/hoastar/orange/tools/app"
)

func TemplateList(c *gin.Context) {
	var (
		err 			error
		templateList	[]*struct {
			process.TplInfo
			CreateUser string `json:"create_user"`
			CreateName string `json:"create_name"`
		}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	db := orm.Eloquent.Model(&process.TplInfo{}).Joins("" +
		"left join sys_user on sys_user.user_id = p_tpl_info.creator").
		Select("p_tpl_info.id, p_tpl_info.create_time, p_tpl_info.update_time, p_tpl_info.`name`, p_tpl_info.`creator`, " +
			"sys_user.username as create_user, sys_user.nick_name as create_name")

	result, err := pagination.Paging(&pagination.Param{
		C: c,
		DB: db,
	}, &templateList, SearchParams, "p_tpl_info")

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, result, "获取模板列表成功")
}

// 创建模板
func CreateTemplate(c *gin.Context) {
	var (
		err				error
		templateValue	process.TplInfo
		templateCount	int
	)

	err = c.ShouldBind(&templateValue)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 确认模板是否存在
	err = orm.Eloquent.Model(&templateValue).
		Where("name = ?", templateValue.Name).
		Count(&templateCount).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	if templateCount > 0 {
		app.Error(c, -1, errors.New("模板名称已存在"), "")
		return
	}

	templateValue.Creator = tools.GetUserId(c)		// 当前登陆用户ID
	err = orm.Eloquent.Create(&templateValue).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.Ok(c, "", "创建成功")
}

// 模板详情
func TemplateDetails(c *gin.Context) {
	var (
		err						error
		templateDetailsValue	process.TplInfo
	)

	templateId := c.DefaultQuery("template_id", "")
	if templateId == "" {
		app.Error(c, -1, err, "")
		return
	}

	err = orm.Eloquent.Model(&templateDetailsValue).
		Where("id = ?", templateId).
		Find(&templateDetailsValue).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, templateDetailsValue, "")
}

// 更新模板
func UpdateTemplate(c *gin.Context) {
	var (
		err				error
		templateValue	process.TplInfo
	)

	err = c.ShouldBind(&templateValue)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	err = orm.Eloquent.Model(&templateValue).
		Where("id = ?", templateValue.Id).
		Updates(map[string]interface{}{
			"name": templateValue.Name,
			"remarks": templateValue.Remarks,
			"form_structure": templateValue.FormStructure,
	}).Error

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.Ok(c, templateValue, "")
}

// 删除模板
func DeleteTemplate(c *gin.Context) {
	var (
		err error
	)

	templateId := c.DefaultQuery("templateId", "")
	if templateId == "" {
		app.Error(c, -1, errors.New("参数不正确，请确认templateId是否传递"), "")
		return
	}

	err = orm.Eloquent.Delete(process.TplInfo{}, "id = ?", templateId).Error

	if err != nil {
		app.Error(c, -1, err, "")
	}

	app.Ok(c, "", "删除模板成功")
}
