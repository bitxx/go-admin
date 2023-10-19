package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/common"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", common.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
