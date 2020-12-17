/*
@Time : 2020/12/9 下午11:32
@Author : hoastar
@File : tpl
@Software: GoLand
*/

package tpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tpl(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
