package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/biz/user/service"
	"go-admin/internal/app/biz/user/service/dto"
	cLang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/api"
	_ "go-admin/pkg/dto/response"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/dateutils"
	"time"
)

type UserOperLog struct {
	api.Api
}

// GetPage app-获取用户操作日志分页列表
func (e UserOperLog) GetPage(c *gin.Context) {
	req := dto.UserOperLogQueryReq{}
	s := service.UserOperLog{}
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
	req.ShowInfo = false
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Get app-获取用户操作日志详情
func (e UserOperLog) Get(c *gin.Context) {
	req := dto.UserOperLogGetReq{}
	s := service.UserOperLog{}
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

// Export app-导出用户操作日志
func (e UserOperLog) Export(c *gin.Context) {
	req := dto.UserOperLogQueryReq{}
	s := service.UserOperLog{}
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
	req.ShowInfo = true
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.Export(list)
	fileName := "user-oper-log_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
