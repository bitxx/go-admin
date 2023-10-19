package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/constant"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	adminService "go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/dto/api"
	"go-admin/common/middleware"
	"go-admin/common/middleware/auth"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"time"
)

type SysConfig struct {
	api.Api
}

// GetPage 获取配置管理列表
func (e SysConfig) GetPage(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Get 获取配置管理
func (e SysConfig) Get(c *gin.Context) {
	req := dto.SysConfigGetReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Insert 创建配置管理
func (e SysConfig) Insert(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Update 修改配置管理
func (e SysConfig) Update(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.Update(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Delete 删除配置管理
func (e SysConfig) Delete(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Remove(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// GetSysConfigBySysApp 获取系统配置信息
func (e SysConfig) GetSysConfigBySysApp(c *gin.Context) {
	req := dto.SysConfigQueryReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	// 控制只读前台的数据
	list := make([]models.SysConfig, 0)
	req.IsFrontend = constant.SysConfIsFrontend
	list, respCode, err := s.GetWithKeyList(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	mp := make(map[string]string)
	for i := 0; i < len(list); i++ {
		key := list[i].ConfigKey
		if key != "" {
			mp[key] = list[i].ConfigValue
		}
	}
	e.OK(mp, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// GetSysConfigByKey 根据Key获取SysConfig的Service
func (e SysConfig) GetSysConfigByKey(c *gin.Context) {
	var s = new(service.SysConfig)
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.SysConfigByKeyResp)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	resp, respCode, err := s.GetByKey(req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(resp, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Export 导出字典
func (e SysConfig) Export(c *gin.Context) {
	req := dto.SysConfigQueryReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	//最小导出数据量
	maxSize, respCode, err := sysConfService.GetWithKeyInt("sys_max_export_size")
	if err != nil {
		e.Error(respCode, err.Error())
	}

	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.GetExcel(list)
	fileName := "config_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
