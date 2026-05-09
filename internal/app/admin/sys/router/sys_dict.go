package router

import (
	"github.com/gin-gonic/gin"
	apis2 "go-admin/internal/app/admin/sys/apis"

	"go-admin/pkg/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDictRouter)
}

func registerDictRouter(v1 *gin.RouterGroup) {
	dictApi := apis2.SysDictType{}
	dataApi := apis2.SysDictData{}
	dicts := v1.Group("/admin/sys/sys-dict").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{

		dicts.GET("/data", dataApi.GetPage)
		dicts.GET("/data/:id", dataApi.Get)
		dicts.POST("/data", dataApi.Insert)
		dicts.PUT("/data/:id", dataApi.Update)
		dicts.DELETE("/data", dataApi.Delete)
		dicts.GET("/data/select", dataApi.GetList)

		dicts.GET("/type/option-select", dictApi.GetList)
		dicts.GET("/type", dictApi.GetPage)
		dicts.GET("/type/:id", dictApi.Get)
		dicts.POST("/type", dictApi.Insert)
		dicts.GET("/type/export", dictApi.Export)
		dicts.PUT("/type/:id", dictApi.Update)
		dicts.DELETE("/type", dictApi.Delete)
	}
}
