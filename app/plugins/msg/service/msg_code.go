package service

import (
	"go-admin/app/plugins/msg/models"
	"go-admin/app/plugins/msg/service/dto"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get plugins-获取验证码管理详情
func (e *MsgCode) Get(id int64, p *middleware.DataPermission) (*models.MsgCode, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.MsgCode{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// QueryOne plugins-获取验证码管理一条记录
func (e *MsgCode) QueryOne(queryCondition *dto.MsgCodeQueryReq, p *middleware.DataPermission) (*models.MsgCode, int, error) {
	data := &models.MsgCode{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// Count admin-获取验证码管理数据总数
func (e *MsgCode) Count(queryCondition *dto.MsgCodeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.MsgCode{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}
