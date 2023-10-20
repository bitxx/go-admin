package response

import (
	"go-admin/core/utils/strutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code int, msg string) {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(strutils.GenerateMsgIDFromContext(c))
	res.SetCode(int32(code))
	c.Set("result", res)
	c.Set("status", code)
	//status := code
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func OKByCode(c *gin.Context, data interface{}, code int, msg string) {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(strutils.GenerateMsgIDFromContext(c))
	res.SetData(data)
	res.SetCode(int32(code))
	//res.SetSuccess(false) //多余，暂不使用
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	//res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(strutils.GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func Download(c *gin.Context, data []byte, filename, contentType string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Disposition", "attachment;filename="+filename)
	c.Header("Access-Control-Expose-Headers", "Content-Disposition") //允许获取懂啊指定header
	c.Header("Content-AppType", contentType)
	c.Data(http.StatusOK, contentType, data)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result, extend interface{}, count int64, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Extend = extend
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	data["requestId"] = strutils.GenerateMsgIDFromContext(c)
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}
