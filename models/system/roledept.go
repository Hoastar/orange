/*
@Time : 2020/11/1 下午9:06
@Author : hoastar
@File : roledept
@Software: GoLand
*/

package system

import (
	"fmt"
	"github.com/hoastar/orange/global/orm"
)

type SysRoleDept struct {
	RoleId int `gorm:"type:int(11)"`
	DeptId int `gorm:"type:int(11)"`
	Id		int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
func (rm *SysRoleDept) Insert(roleId int, deptIds []int) (bool, error) {
sql := "INSERT INTO `sys_role_dept` (`role_id`,`dept_id`) VALUES "

	for i := 0; i < len(deptIds); i++ {
		if len(deptIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", roleId, deptIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, deptIds[i])
		}
	}
	orm.Eloquent.Exec(sql)

	return true, nil
}

func (rm *SysRoleDept) DeleteRoleDept(roleId int) (bool, error) {
	if err := orm.Eloquent.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}
	var role SysRole
	if err := orm.Eloquent.Table("sys_role").Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}

	return true, nil

}