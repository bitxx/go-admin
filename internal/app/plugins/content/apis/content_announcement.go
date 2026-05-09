package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/plugins/content/service"
	"go-admin/internal/app/plugins/content/service/dto"
	cLang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/api"
	_ "go-admin/pkg/dto/response"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/middleware/auth"
	"go-admin/pkg/utils/dateutils"
	"time"
)

type ContentAnnouncement struct {
	api.Api
}

// GetPage plugins-获取公告管理分页列表
func (e ContentAnnouncement) GetPage(c *gin.Context) {
	req := dto.ContentAnnouncementQueryReq{}
	s := service.ContentAnnouncement{}
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

// Get plugins-获取公告管理详情
func (e ContentAnnouncement) Get(c *gin.Context) {
	req := dto.ContentAnnouncementGetReq{}
	s := service.ContentAnnouncement{}
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

// Insert plugins-新增公告管理
func (e ContentAnnouncement) Insert(c *gin.Context) {
	req := dto.ContentAnnouncementInsertReq{}
	s := service.ContentAnnouncement{}
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

// Update plugins-更新公告管理
func (e ContentAnnouncement) Update(c *gin.Context) {
	req := dto.ContentAnnouncementUpdateReq{}
	s := service.ContentAnnouncement{}
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

// Delete plugins-删除公告管理
func (e ContentAnnouncement) Delete(c *gin.Context) {
	s := service.ContentAnnouncement{}
	req := dto.ContentAnnouncementDeleteReq{}
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

// Export plugins-导出公告管理
func (e ContentAnnouncement) Export(c *gin.Context) {
	req := dto.ContentAnnouncementQueryReq{}
	s := service.ContentAnnouncement{}
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
	fileName := "content-announcement_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
