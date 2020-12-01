/*
@Time : 2020/11/26 下午3:52
@Author : hoastar
@File : classify
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

// 创建流程分类
func CreateClassify(c *gin.Context) {
	var (
		err				error
		classifyValue	process.Classify
		classifyCount	int
	)

	err = c.ShouldBind(&classifyValue)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 判断创建的流程是否存在
	err = orm.Eloquent.Table("p_process_classify").
		Where("name = ?", classifyValue.Name).
		Where("`delete_time` IS NULL").
		Count(&classifyCount).Error

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	if classifyCount > 0 {
		app.Error(c, -1, errors.New("创建的分类名称已经存在"), "")
		return
	}
	classifyValue.Creator = tools.GetUserId(c)

	err = orm.Eloquent.Table("p_process_classify").Create(&classifyValue).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, "", "创建流程成功")
}

// 流程分类列表
func ClassifyList(c *gin.Context) {
	type classifyValue struct {
		process.Classify
		CreateUser string `json:"create_user"`
		CreateName string `json:"create_name"`
	}

	var (
		err 			error
		classifyList	[]*classifyValue
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	db := orm.Eloquent.Model(&process.Classify{}).Joins("" +
		"left join sys_user on sys_user.user_id = p_process_classify.creator").
		Select("p_process_classify.*, sys_user.username as create_user, sys_user.nick_name as create_name").
		Where("p_process_classify.`delete_time` IS NULL")

	result, err := pagination.Paging(&pagination.Param{
		C: c,
		DB: db,
	}, &classifyList, SearchParams, "p_process_classify")

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.Ok(c, result, "获取分类列表成功")
}

// 更新流程分类

func UpdateClassify(c *gin.Context) {
	var (
		err  			error
		classifyValue		process.Classify
	)

	err = c.ShouldBind(&classifyValue)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	// 更新
	err = orm.Eloquent.Model(&classifyValue).
		Where("id = ?", classifyValue.Id).
		Update("name", classifyValue.Name).Error

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, classifyValue, "流程分类更新成功")
}

// 删除流程分类
func DeleteClassify(c *gin.Context) {
	classifyId := c.DefaultQuery("classifyId", "")
	if classifyId == "" {
		app.Error(c, -1, errors.New("参数传递失败，请确认classifyId是否传递"), "")
		return
	}

	err := orm.Eloquent.Delete(process.Classify{}, "id = ?", classifyId).Error
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.Ok(c, "", "流程分类删除成功")
}