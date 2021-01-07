/*
@Time : 2020/12/20 下午3:40
@Author : hoastar
@File : connection
@Software: GoLand
*/

package ldap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/spf13/viper"
	"time"
)

var conn *ldap.Conn

// ldap连接
func ldapConnection() (err error) {
	var ldapConn = fmt.Sprintf("%v:%v", viper.GetString("settings.ldap.host"), viper.GetString("settings.ldap.port"))

	if viper.GetBool("settings.ldap.tls") {
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		conn, err = ldap.DialTLS("tcp", ldapConn, tlsConf)
	} else {
		conn, err = ldap.Dial("tcp", ldapConn)
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("无法连接到ldap服务器，%v", err))
		logger.Error(err)
		return
	}

	//设置超时时间
	conn.SetTimeout(5 * time.Second)

	return
}