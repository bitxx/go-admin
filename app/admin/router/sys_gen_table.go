package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysTableRouter)
}

func registerSysTableRouter(v1 *gin.RouterGroup) {
	api := apis.SysTables{}
	r := v1.Group("/sys/table").Use(middleware.Auth())
	{
		r.GET("", api.GetPage)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/:id", api.Get)
		r.GET("/dbtables", api.GetDBTablePage) //获取当前数据库表

		//代码生成相关
		r.GET("/preview/:id", api.Preview)           //预览
		r.GET("/gen/:id", api.GenCode)               //生成代码
		r.GET("/gen/download/:id", api.DownloadCode) //下载代码
		r.GET("/gen/db/:id", api.GenDB)              //生成菜单
	}
}
