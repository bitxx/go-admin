package service

import (
	"errors"
	"go-admin/internal/app/plugins/msg/models"
	"go-admin/internal/app/plugins/msg/service/dto"
	cLang "go-admin/internal/common/lang"
	cDto "go-admin/pkg/dto"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"gorm.io/gorm"
)

type MsgCode struct {
	service.Service
}

// NewMsgCodeService plugins-实例化验证码管理
func NewMsgCodeService(s *service.Service) *MsgCode {
	var srv = new(MsgCode)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage plugins-获取验证码管理分页列表
func (e *MsgCode) GetPage(c *dto.MsgCodeQueryReq, p *middleware.DataPermission) ([]models.MsgCode, int64, int, error) {
	var data models.MsgCode
	var list []models.MsgCode
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, cLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, cLang.DataQueryCode, cLang.DataQueryLogCode, err)
	}
	return list, count, cLang.SuccessCode, nil
}

// Get plugins-获取验证码管理详情
func (e *MsgCode) Get(id int64, p *middleware.DataPermission) (*models.MsgCode, int, error) {
	if id <= 0 {
		return nil, cLang.ParamErrCode, lang.MsgErr(cLang.ParamErrCode, e.Lang)
	}
	data := &models.MsgCode{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, cLang.DataQueryCode, cLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cLang.DataNotFoundCode, lang.MsgErr(cLang.DataNotFoundCode, e.Lang)
	}
	return data, cLang.SuccessCode, nil
}

// QueryOne plugins-获取验证码管理一条记录
func (e *MsgCode) QueryOne(queryCondition *dto.MsgCodeQueryReq, p *middleware.DataPermission) (*models.MsgCode, int, error) {
	data := &models.MsgCode{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, cLang.DataQueryCode, cLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cLang.DataNotFoundCode, lang.MsgErr(cLang.DataNotFoundCode, e.Lang)
	}
	return data, cLang.SuccessCode, nil
}

// Count admin-获取验证码管理数据总数
func (e *MsgCode) Count(queryCondition *dto.MsgCodeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.MsgCode{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, cLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, cLang.DataQueryCode, cLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, cLang.DataNotFoundCode, lang.MsgErr(cLang.DataNotFoundCode, e.Lang)
	}
	return count, cLang.SuccessCode, nil
}
