/*
@Time : 2020/10/28 下午10:41
@Author : hoastar
@File : interface
@Software: GoLand
*/

package database

import "github.com/jinzhu/gorm"

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
	GetConnect() string
}
