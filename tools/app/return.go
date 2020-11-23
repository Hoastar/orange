/*
@Time : 2020/11/5 下午11:52
@Author : hoastar
@File : return
@Software: GoLand
*/

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/pkg/logger"
	"net/http"
)

func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()

	if msg != "" {
		res.Msg = msg
	}
	logger.Error(res.Msg)
	c.JSON(http.StatusOK, res.ReturnError(code))
}

func Ok(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}

// 分页处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}









































