package service

import (
	"go-admin/app/plugins/msg/models"
	"go-admin/app/plugins/msg/service/dto"
	cDto "go-admin/common/dto"
	"go-admin/common/dto/service"
	"go-admin/common/middleware"
	"go-admin/config/lang"
	"gorm.io/gorm"
)

type MsgCode struct {
	service.Service
}

// NewMsgCodeService
// @Description: 实例化MsgCode
// @param s
// @return *MsgCode
func NewMsgCodeService(s *service.Service) *MsgCode {
	var srv = new(MsgCode)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取MsgCode列表
// @receiver e
// @param c
// @param p
// @return []models.MsgCode
// @return int64
// @return int
// @return error
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

// Get
// @Description: 获取MsgCode对象
// @receiver e
// @param id 编号
// @param p
// @return *models.MsgCode
// @return int
// @return error
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

// QueryOne
// @Description: 通过自定义条件获取MsgCode一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.MsgCode
// @return error
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

// Count
//
//	@Description: 获取条数
//	@receiver e
//	@param c
//	@return int64
//	@return int
//	@return error
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
