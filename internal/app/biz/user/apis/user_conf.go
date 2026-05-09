package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/app/biz/user/service"
	"go-admin/internal/app/biz/user/service/dto"
	cLang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/api"
	_ "go-admin/pkg/dto/response"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/middleware/auth"
)

type UserConf struct {
	api.Api
}

// GetPage app-获取用户配置管理分页列表
func (e UserConf) GetPage(c *gin.Context) {
	req := dto.UserConfQueryReq{}
	s := service.UserConf{}
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

// Get app-获取用户配置管理详情
func (e UserConf) Get(c *gin.Context) {
	req := dto.UserConfGetReq{}
	s := service.UserConf{}
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

// Update app-更新用户配置管理
func (e UserConf) Update(c *gin.Context) {
	req := dto.UserConfUpdateReq{}
	s := service.UserConf{}
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
