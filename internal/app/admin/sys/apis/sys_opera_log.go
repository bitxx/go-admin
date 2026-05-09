package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/admin/sys/service/dto"
	cLang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/api"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/dateutils"
	"time"
)

type SysOperLog struct {
	api.Api
}

// GetPage admin-获取操作日志分页列表
func (e SysOperLog) GetPage(c *gin.Context) {
	s := adminService.SysOperLog{}
	req := dto.SysOperLogQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(cLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, cLang.DataDecodeCode, cLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Get admin-获取操作日志详情
func (e SysOperLog) Get(c *gin.Context) {
	s := new(adminService.SysOperLog)
	req := dto.SysOperLogGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(cLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, cLang.DataDecodeCode, cLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Delete admin-删除操作日志
func (e SysOperLog) Delete(c *gin.Context) {
	s := new(adminService.SysOperLog)
	req := dto.SysOperLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(cLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, cLang.DataDecodeCode, cLang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Delete(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Export admin-导出操作日志
func (e SysOperLog) Export(c *gin.Context) {
	req := dto.SysOperLogQueryReq{}
	s := adminService.SysOperLog{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(cLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, cLang.DataDecodeCode, cLang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	maxSize, respCode, err := sysConfService.GetWithKeyInt("admin_sys_max_export_size")
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
	data, _ := s.Export(list)
	fileName := "operlog_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
