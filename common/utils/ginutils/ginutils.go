package ginutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler http.Handler 转换成 gin.HandlerFunc
func Handler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
