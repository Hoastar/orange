/*
@Time : 2020/11/3 下午10:09
@Author : hoastar
@File : requestid
@Software: GoLand
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming(request) header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4 := uuid.New()
			requestId = u4.String()
		}

		// Set is used to store a new key/value. Expose it for use in the application，
		c.Set("X-Request-Id", requestId)

		// Write to header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
