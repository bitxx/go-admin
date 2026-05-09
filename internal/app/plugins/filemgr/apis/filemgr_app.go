package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/plugins/filemgr/service"
	"go-admin/internal/app/plugins/filemgr/service/dto"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/api"
	_ "go-admin/pkg/dto/response"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/middleware/auth"
	"go-admin/pkg/utils/dateutils"
	"mime/multipart"
	"time"
)

type FilemgrApp struct {
	api.Api
}

// GetPage plugins-获取APP管理分页列表
func (e FilemgrApp) GetPage(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(clang.SuccessCode, e.Lang))
}

// Get plugins-获取APP管理详情
func (e FilemgrApp) Get(c *gin.Context) {
	req := dto.FilemgrAppGetReq{}
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(clang.SuccessCode, e.Lang))
}

// Insert plugins-新增APP管理
func (e FilemgrApp) Insert(c *gin.Context) {
	req := dto.FilemgrAppInsertReq{}
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
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
	e.OK(id, lang.MsgByCode(clang.SuccessCode, e.Lang))
}

// Delete plugins-删除APP管理
func (e FilemgrApp) Delete(c *gin.Context) {
	s := service.FilemgrApp{}
	req := dto.FilemgrAppDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Delete(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(clang.SuccessCode, e.Lang))
}

// Upload  plugins-上传APP
func (e FilemgrApp) Upload(c *gin.Context) {
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
		return
	}
	form, err := e.Context.MultipartForm()
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
		return
	}

	//获取上传文件信息
	var filePath string
	file := &multipart.FileHeader{}

	respCode, err := s.GetSingleUploadFileInfo(form, file, &filePath)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}

	//保存上传文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		e.Error(clang.AppUploadLogCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.AppUploadCode, clang.AppUploadLogCode, err).Error())
		return
	}
	e.OK(filePath, lang.MsgByCode(clang.AppUploadSuccessCode, e.Lang))
}

// Update  plugins-更新APP管理
func (e FilemgrApp) Update(c *gin.Context) {
	req := dto.FilemgrAppUpdateReq{}
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
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
		e.OK(nil, lang.MsgByCode(clang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(clang.SuccessCode, e.Lang))
}

// Export  plugins-导出APP管理
func (e FilemgrApp) Export(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(clang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, clang.DataDecodeCode, clang.DataDecodeLogCode, err).Error())
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
	fileName := "filemgr-app_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
