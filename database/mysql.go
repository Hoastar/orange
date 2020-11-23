/*
@Time : 2020/10/28 下午10:32
@Author : hoastar
@File : mysql
@Software: GoLand
*/

package database

import (
	"bytes"
	"strconv"

	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/hoastar/orange/tools/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Defined database info
var (
	Dbtype   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

type Mysql struct {
}

func (m *Mysql) Setup() {
	var err error
	var db Database

	db = new(Mysql)
	orm.MysqlConn = db.GetConnect()
	orm.Eloquent, err = db.Open(Dbtype, orm.MysqlConn)

	if err != nil {
		logger.Fatalf("%s connect error %v", Dbtype, err)
	} else {
		logger.Infof("%s connect success!", Dbtype)
	}

	if orm.Eloquent.Error != nil {
		logger.Fatalf("database error %v", orm.Eloquent.Error)
	}

	// 是否开启详细日志记录
	orm.Eloquent.LogMode(viper.GetBool("settings.gorm.logMode"))

	// set maxOpenConn
	orm.Eloquent.DB().SetMaxOpenConns(viper.GetInt("settings.gorm.maxopenconn"))

	// set maxIdleConn
	orm.Eloquent.DB().SetMaxIdleConns(viper.GetInt("settings.gorm.maxidleconn"))
}

// Open initialize a new db connection
func (m *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

func (m *Mysql) GetConnect() string {
	Dbtype = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}
