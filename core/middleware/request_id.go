package middleware

import (
	"go-admin/core/global"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequestId 自动增加requestId
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		requestId := c.GetHeader(global.TrafficKey)
		if requestId == "" {
			requestId = c.GetHeader(strings.ToLower(global.TrafficKey))
		}
		if requestId == "" {
			requestId = idgen.UUID()
		}
		c.Request.Header.Set(global.TrafficKey, requestId)
		c.Set(global.TrafficKey, requestId)
		c.Set(global.LoggerKey, log.WithFields(map[string]interface{}{global.TrafficKey: requestId}))
		c.Next()
	}
}
