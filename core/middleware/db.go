package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/core/runtime"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", runtime.RuntimeConfig.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
