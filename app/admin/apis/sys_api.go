package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	adminService "go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/dto/api"
	_ "go-admin/common/dto/response"
	"go-admin/common/middleware"
	"go-admin/common/middleware/auth"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"time"
)

type SysApi struct {
	api.Api
}

// GetPage 获取接口管理列表
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiQueryReq{}
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

// Get 获取接口管理
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
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

// Update 修改接口管理
// @Summary 修改接口管理
// @Description 修改接口管理
// @Tags 接口管理
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
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

// Delete 删除接口管理
func (e SysApi) Delete(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
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

// Export 导出接口
func (e SysApi) Export(c *gin.Context) {
	req := dto.SysApiQueryReq{}
	s := service.SysApi{}
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
	fileName := "api_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
