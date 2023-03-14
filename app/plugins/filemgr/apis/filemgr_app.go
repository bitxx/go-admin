package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/service"
	fLang "go-admin/app/plugins/filemgr/lang"
	"go-admin/app/plugins/filemgr/service"
	"go-admin/app/plugins/filemgr/service/dto"
	"go-admin/common/core/api"
	_ "go-admin/common/core/pkg/response"
	"go-admin/common/middleware"
	"go-admin/common/middleware/auth"
	"go-admin/common/utils/dateUtils"
	"go-admin/config/lang"
	"mime/multipart"
	"time"
)

type FilemgrApp struct {
	api.Api
}

// GetPage
// @Description: 获取App管理列表
// @receiver e
// @param c
func (e FilemgrApp) GetPage(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
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

// Get
// @Description: 获取App管理
// @receiver e
// @param c
func (e FilemgrApp) Get(c *gin.Context) {
	req := dto.FilemgrAppGetReq{}
	s := service.FilemgrApp{}
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

// Insert
// @Description: 创建App管理
// @receiver e
// @param c
func (e FilemgrApp) Insert(c *gin.Context) {
	req := dto.FilemgrAppInsertReq{}
	s := service.FilemgrApp{}
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

// Delete
// @Description:App管理
// @receiver e
// @param c
func (e FilemgrApp) Delete(c *gin.Context) {
	s := service.FilemgrApp{}
	req := dto.FilemgrAppDeleteReq{}
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

// Upload 上传App文件
// @Summary 上传App文件
// @Description 上传App文件
// @Tags App管理
// @Router /api/v1/app-manager/upload [post]
func (e FilemgrApp) Upload(c *gin.Context) {
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	form, err := e.Context.MultipartForm()
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
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
		e.Error(fLang.PluginsAppUploadLogCode, lang.MsgLogErrf(e.Logger, e.Lang, fLang.PluginsAppUploadCode, fLang.PluginsAppUploadLogCode, err).Error())
		return
	}
	e.OK(filePath, lang.MsgByCode(fLang.PluginsAppUploadSuccessCode, e.Lang))
}

// Update
// @Description: 修改App管理
// @receiver e
// @param c
func (e FilemgrApp) Update(c *gin.Context) {
	req := dto.FilemgrAppUpdateReq{}
	s := service.FilemgrApp{}
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

// Export
// @Description: 导出App管理
// @receiver e
// @param c
func (e FilemgrApp) Export(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
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
	fileName := "filemgr-app_" + dateUtils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
