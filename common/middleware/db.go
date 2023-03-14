package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/core"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", core.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
