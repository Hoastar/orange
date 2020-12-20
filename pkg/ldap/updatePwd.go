/*
@Time : 2020/12/20 下午3:44
@Author : hoastar
@File : updatePwd
@Software: GoLand
*/

package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/unicode"
)

func LdapUpdatePwd(username string, oldPassword string, newPassword string) (err error) {
	err = ldapConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	var userDn = fmt.Sprintf("cn=%v,%v", username, viper.GetString("settings.ldap.baseDn"))

	err = conn.Bind(userDn, oldPassword)
	if err != nil {
		logger.Error("用户或密码错误。", err)
		return
	}

	sql2 := ldap.NewModifyRequest(userDn, nil)

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, _ := utf16.NewEncoder().String(newPassword)

	sql2.Replace("unicodePwd", []string{pwdEncoded})
	sql2.Replace("userAccountControl", []string{"512"})

	if err = conn.Modify(sql2); err != nil {
		logger.Error("密码修改失败，%v", err.Error())
		return
	}

	return
}