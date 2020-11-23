/*
@Time : 2020/11/5 下午10:53
@Author : hoastar
@File : user
@Software: GoLand
*/

package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
	"github.com/hoastar/orange/pkg/logger"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}
	return claims.(jwt.MapClaims)

}

func GetUserName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["nice"] != nil {
		return (data["nice"].(string))
	}
	fmt.Println("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  缺少nice")
	return ""
}

func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int((data["identity"]).(float64))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  说明：缺少identity")
	return 0
}

func GetUserIdToStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return Int64ToString(int64((data["identity"]).(float64)))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  缺少identity")
	return ""
}