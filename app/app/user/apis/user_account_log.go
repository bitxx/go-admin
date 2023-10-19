package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/service"
	"go-admin/app/app/user/service"
	"go-admin/app/app/user/service/dto"
	"go-admin/common/core/api"
	_ "go-admin/common/core/response"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"time"
)

type UserAccountLog struct {
	api.Api
}

// GetPage
// @Description: 获取账变记录列表
// @receiver e
// @param c
func (e UserAccountLog) GetPage(c *gin.Context) {
	req := dto.UserAccountLogQueryReq{}
	s := service.UserAccountLog{}
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
	req.ShowInfo = false
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Get
// @Description: 获取账变记录
// @receiver e
// @param c
func (e UserAccountLog) Get(c *gin.Context) {
	req := dto.UserAccountLogGetReq{}
	s := service.UserAccountLog{}
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

// Export
// @Description: 导出账变记录
// @receiver e
// @param c
func (e UserAccountLog) Export(c *gin.Context) {
	req := dto.UserAccountLogQueryReq{}
	s := service.UserAccountLog{}
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
	req.ShowInfo = true
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.GetExcel(list)
	fileName := "user-account-log_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
