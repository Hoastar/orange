/*
@Time : 2020/11/14 上午12:58
@Author : hoastar
@File : url
@Software: GoLand
*/

package tools

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 获取url中目批量目标资源的ID
func IdsStrToIdsIntGroup(key string, c *gin.Context) []int {
	return idsStrToIdsIntGroup(c.Param(key))
}

func idsStrToIdsIntGroup(keys string) []int {
	IDS := make([]int, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := StringToInt(ids[i])
		IDS = append(IDS, ID)
	}
	return IDS
}
