/*
@Time : 2020/10/28 下午10:20
@Author : hoastar
@File : initialize
@Software: GoLand
*/

package database

import "github.com/hoastar/orange/tools/config"

func Setup() {
	dbType := config.DatabaseConfig.Dbtype
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}
}