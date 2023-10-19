package middleware

import (
	"go-admin/common/global"
	"go-admin/common/utils/idgen"
	"go-admin/config/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequestId 自动增加requestId
func RequestId(trafficKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		requestId := c.GetHeader(trafficKey)
		if requestId == "" {
			requestId = c.GetHeader(strings.ToLower(trafficKey))
		}
		if requestId == "" {
			requestId = idgen.UUID()
		}
		c.Request.Header.Set(trafficKey, requestId)
		c.Set(trafficKey, requestId)
		c.Set(global.LoggerKey, config.LoggerConfig.GetLogger(trafficKey, requestId))
		c.Next()
	}
}
