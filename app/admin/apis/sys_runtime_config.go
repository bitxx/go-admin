package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	"go-admin/common/core/api"
)

type SysRuntimeConfig struct {
	api.Api
}

// GetPage 获取配置管理列表
func (e SysRuntimeConfig) GetConfig(c *gin.Context) {
	s := service.SysRuntimeConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	config, err := s.GetConfig()
	if err != nil {
		e.Error(500, "获取失败")
		return
	}
	e.OK(config, "获取成功")
}
