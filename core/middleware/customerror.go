package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					c.Status(statusCode)
					log.Error(
						time.Now().Format("2006-01-02 15:04:05"),
						"[ERROR]",
						c.Request.Method,
						c.Request.URL,
						statusCode,
						c.Request.RequestURI,
						iputils.GetClientIP(c),
						p[2],
					)
					c.JSON(http.StatusOK, gin.H{
						"code": statusCode,
						"msg":  p[2],
					})
				}
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
