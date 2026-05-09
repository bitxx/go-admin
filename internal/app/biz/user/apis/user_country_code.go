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
	"go-admin/pkg/middleware/auth"
	"go-admin/pkg/utils/dateutils"
	"time"
)

type UserCountryCode struct {
	api.Api
}

// GetPage biz-获取国家区号管理分页列表
func (e UserCountryCode) GetPage(c *gin.Context) {
	req := dto.UserCountryCodeQueryReq{}
	s := service.UserCountryCode{}
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

// Get biz-获取国家区号管理详情
func (e UserCountryCode) Get(c *gin.Context) {
	req := dto.UserCountryCodeGetReq{}
	s := service.UserCountryCode{}
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

// Insert biz-新增国家区号管理
func (e UserCountryCode) Insert(c *gin.Context) {
	req := dto.UserCountryCodeInsertReq{}
	s := service.UserCountryCode{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(cLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, cLang.DataDecodeCode, cLang.DataDecodeLogCode, err).Error())
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
	e.OK(id, lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Update biz-更新国家区号管理
func (e UserCountryCode) Update(c *gin.Context) {
	req := dto.UserCountryCodeUpdateReq{}
	s := service.UserCountryCode{}
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
		e.OK(nil, lang.MsgByCode(cLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(cLang.SuccessCode, e.Lang))
}

// Delete biz-删除国家区号管理
func (e UserCountryCode) Delete(c *gin.Context) {
	s := service.UserCountryCode{}
	req := dto.UserCountryCodeDeleteReq{}
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

// Export biz-导出国家区号管理
func (e UserCountryCode) Export(c *gin.Context) {
	req := dto.UserCountryCodeQueryReq{}
	s := service.UserCountryCode{}
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
	fileName := "user-country-code_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
