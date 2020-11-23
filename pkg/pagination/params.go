/*
@Time : 2020/11/7 下午11:50
@Author : hoastar
@File : params
@Software: GoLand
*/

package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/pkg/logger"
)

func RequestParams(c *gin.Context) map[string]interface{} {
	params := make(map[string]interface{}, 10)

	if c.Request.Form == nil {
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			logger.Error(err)
		}
	}

	if len(c.Request.Form) > 0 {
		for key, value := range c.Request.Form {
			if key == "page" || key == "per_page" || key == "sort" {
				continue
			}
			params[key] = value[0]

		}
	}
	return params

}
