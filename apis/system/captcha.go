/*
@Time : 2020/11/15 下午2:43
@Author : hoastar
@File : captcha
@Software: GoLand
*/

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/tools/app"
	"github.com/hoastar/orange/tools/captcha"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("验证码获取失败, %v", err.Error()))
		return
	}
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
