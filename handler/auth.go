/*
@Time : 2020/11/3 下午11:37
@Author : hoastar
@File : auth
@Software: GoLand
*/

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/models/system"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/hoastar/orange/tools"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

// PayloadFunc 添加额外业务相关的信息
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)

		return jwt.MapClaims{
			jwt.IdentityKey: u.UserId,
			jwt.RoleIdKey: r.RoleId,
			jwt.RoleKey: r.RoleKey,
			jwt.NiceKey:  u.Username,
			jwt.RoleNameKey: r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

// IdentityHandler
func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{} {
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
	}
}

// Authenticator 在登录接口中使用的验证方法，并返回验证成功后的用户对象
// @Summary 登陆
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Param username body models.Login  true "Add account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var (
		//err           error
		loginVal      system.Login
		loginLog      system.LoginLog
		//roleValue     system.SysRole
		//authUserCount int
		//addUserInfo   system.SysUser
		//ldapUserInfo  *ldap.Entry
	)

	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()
	loginLog.Ipaddr = c.ClientIP()
	loginLog.LoginLocation = tools.GetLocation(c.ClientIP())
	loginLog.Status = "0"
	loginLog.Remark = c.Request.UserAgent()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Msg = "登录成功"
	loginLog.Platform = ua.Platform()

	// 获取前端传过来的表单数据
	if err := c.ShouldBind(&loginVal); err != nil {
		loginLog.Status = "1"
		loginLog.Msg = "解析数据失败"
		loginLog.Username = loginVal.Username
		_, _ = loginLog.Create()
		return nil, jwt.ErrInvalidVerificationode
	}
	loginLog.Username = loginVal.Username

	// store.Verify 校验验证码
	if !store.Verify(loginVal.UUID, loginVal.Code, true) {
	loginLog.Status = "1"
	loginLog.Msg = "验证码错误"
	_, _ = loginLog.Create()
	return nil, jwt.ErrInvalidVerificationode
	}

	// ldap 验证 待理解后补充


	user, role, e := loginVal.GetUser()
	if e == nil {
		_, _ = loginLog.Create()
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		loginLog.Msg = "登录失败"
		loginLog.Status = "1"
		loginLog.Create()
		logger.Info(e.Error())
	}

	return nil, jwt.ErrFailedAuthentication

}

// @Summary 退出登录
// @Descrption 获取token
// Logout can be used by clients to exit system.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security

func LogOut(c *gin.Context) {
	var loginlog system.LoginLog
	ua := user_agent.New(c.Request.UserAgent())
	loginlog.Ipaddr = c.ClientIP()
	location := tools.GetLocation(c.ClientIP())
	loginlog.LoginLocation = location
	loginlog.Status = "0"
	loginlog.Remark = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginlog.Browser = browserName + " " + browserVersion
	loginlog.Os = ua.OS()
	loginlog.Platform = ua.Platform()
	loginlog.Username = tools.GetUserName(c)
	loginlog.Msg = "退出成功"
	_, _ = loginlog.Create()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "退出成功",
	})
}

// Authorizator 登录后其他接口验证传入登录信息的方法
func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)

		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.UserName)
		return true
	}

	return false
}


// Unauthorized 验证失败后设置错误信息
func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": message,
	})
}
