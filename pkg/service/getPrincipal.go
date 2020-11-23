/*
@Time : 2020/11/8 下午8:49
@Author : hoastar
@File : getPrincipal
@Software: GoLand
*/

package service

import (
	"errors"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/system"
	"reflect"
	"strings"
)

func GetPrincipal(processor []int, processMethod string) (principals string, err error) {
	/*
	person 个人
	persongroup 人员组
	department 部门
	variable 变量
	 */

	var principalList []string
	switch processMethod {
	case "person":
		err = orm.Eloquent.Model(&system.SysUser{}).
			Where("user_id in (?)", processor).
			Pluck("nick_name", &principalList).Error

		if err != nil {
			return
		}

	case "variable":
		for _, p := range processor {
			switch p {
			case 1:
				principalList = append(principalList, "创建者")
			case 2:
				principalList = append(principalList, "创建这负责人")
			}
		}
	}
	return strings.Join(principalList, ","), nil
}

func GetPrincipalUserInfo(stateList []interface{}, cretor int) (userInfoList []system.SysUser, err error) {
	var (
		userInfo 			system.SysUser
		deptInfo			system.Dept
		// 用于保持临时查询的列表数据
		userInfoListTmp		[]system.SysUser
		processorList		[]interface{}
	)

	err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", cretor).Find(&userInfo).Error
	if err != nil {
		return
	}


	for _, stateItem := range stateList {
		if reflect.TypeOf(stateItem.(map[string]interface{})["processor"]) == nil {
			err = errors.New("未找到对应的处理人，请确认。")
			return
		}

		stateItemType := reflect.TypeOf(stateItem.(map[string]interface{})["process"]).String()
		if stateItemType == "[]int" {
			for _, v := range stateItem.(map[string]interface{})["processor"].([]int) {
				processorList = append(processorList, v)
			}
		} else {
			processorList = stateItem.(map[string]interface{})["processor"].([]interface{})
		}

		switch stateItem.(map[string]interface{})["process_method"] {
		case "person":
			err = orm.Eloquent.Model(&system.SysUser{}).
				Where("user_id int (?)", processorList).Find(&userInfoListTmp).Error
			if err != nil {
				return
			}

			userInfoList = append(userInfoList, userInfoListTmp...)
		case "variable": // 变量
		for _, processor := range processorList {
			if int(processor.(float64)) == 1 {
				// 创建者
				userInfoList = append(userInfoList, userInfo)
			} else if int(processor.(float64)) == 2 {
				// 1. 查询部门信息
				err = orm.Eloquent.Model(&deptInfo).Where("dept_id = ?", userInfo.DeptId).Find(&deptInfo).Error
				if err != nil {
					return
				}
				// 查询Leader信息
				err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", deptInfo.Leader).Find(&userInfo).Error
				if err != nil {
					return
				}

				userInfoList = append(userInfoList, userInfo)
			}
		}
		}
	}
	return
}
































