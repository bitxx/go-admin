package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"

	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDictRouter)
}

func registerDictRouter(v1 *gin.RouterGroup) {
	dictApi := apis.SysDictType{}
	dataApi := apis.SysDictData{}
	dicts := v1.Group("/dict").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{

		dicts.GET("/data", dataApi.GetPage)
		dicts.GET("/data/:id", dataApi.Get)
		dicts.POST("/data", dataApi.Insert)
		dicts.PUT("/data/:id", dataApi.Update)
		dicts.GET("/data/export", dataApi.Export)
		dicts.DELETE("/data", dataApi.Delete)

		dicts.GET("/type-option-select", dictApi.GetAll)
		dicts.GET("/type", dictApi.GetPage)
		dicts.GET("/type/:id", dictApi.Get)
		dicts.POST("/type", dictApi.Insert)
		dicts.GET("/type/export", dictApi.Export)
		dicts.PUT("/type/:id", dictApi.Update)
		dicts.DELETE("/type", dictApi.Delete)
	}
	opSelect := v1.Group("/dict-data").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		opSelect.GET("/option-select", dataApi.GetSysDictDataAll)
	}
}
