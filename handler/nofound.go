/*
@Time : 2020/11/3 下午11:33
@Author : hoastar
@File : nofound
@Software: GoLand
*/

package handler

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/hoastar/orange/pkg/jwtauth"
	"github.com/hoastar/orange/pkg/logger"
	"net/http"
)

func NoFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	logger.Infof("NoRoute claims: %#v\n", claims)

	c.JSON(http.StatusOK, gin.H{
		"code": "404",
		"message": "not found",
	})
}
