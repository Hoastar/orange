/*
@Time : 2020/10/28 下午10:25
@Author : hoastar
@File : env
@Software: GoLand
*/

package tools

type (
	Mode string
)

const (
	ModeDev Mode = "dev"	// 开发模式
	ModeTest Mode = "test"	// 测试模式
	ModeProd Mode = "prod"	// 生产模式
	Mysql = "mysql"			// mysql数据库类型标识
	Sqlite = "sqlite"		// sqllite
)
