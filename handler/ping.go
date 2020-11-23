/*
@Time : 2020/11/3 下午11:32
@Author : hoastar
@File : ping
@Software: GoLand
*/

package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
