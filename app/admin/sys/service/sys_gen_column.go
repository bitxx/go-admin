package service

import (
	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm"
	"time"
)

type SysGenColumn struct {
	service.Service
}

// NewSysColumnsService admin-实例化表字段管理
func NewSysColumnsService(s *service.Service) *SysGenColumn {
	var srv = new(SysGenColumn)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetList admin-获取表字段全部列表
func (e *SysGenColumn) GetList(c *dto.SysGenColumnQueryReq, p *middleware.DataPermission) ([]models.SysGenColumn, int, error) {
	var list []models.SysGenColumn
	var data models.SysGenColumn
	var count int64

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, baseLang.SuccessCode, nil
}

// Get admin-获取表字段详情
func (e *SysGenColumn) Get(id int64, p *middleware.DataPermission) (*models.SysGenColumn, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysGenColumn{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Insert admin-新增表字段
func (e *SysGenColumn) Insert(c dto.SysGenColumnInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	now := time.Now()
	data := models.SysGenColumn{}
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update admin-更新表字段
func (e *SysGenColumn) Update(c *dto.SysGenColumnUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}

	if c.ColumnComment != "" && data.ColumnComment != c.ColumnComment {
		updates["column_comment"] = c.ColumnComment
	}
	if c.GoType != "" && data.GoType != c.GoType {
		updates["go_type"] = c.GoType
	}
	if c.GoField != "" && data.GoField != c.GoField {
		updates["go_field"] = c.GoField
	}
	if c.JsonField != "" && data.JsonField != c.JsonField {
		updates["json_field"] = c.JsonField
	}
	if c.IsRequired != "" && data.IsRequired != c.IsRequired {
		updates["is_required"] = c.IsRequired
	}
	if c.IsQuery != "" && data.IsQuery != c.IsQuery {
		updates["is_query"] = c.IsQuery
	}
	if c.IsList != "" && data.IsList != c.IsList {
		updates["is_list"] = c.IsList
	}
	if c.QueryType != "" && data.QueryType != c.QueryType {
		updates["query_type"] = c.QueryType
	}
	if c.HtmlType != "" && data.HtmlType != c.HtmlType {
		updates["html_type"] = c.HtmlType
	}
	if c.DictType != "" && data.DictType != c.DictType {
		updates["dict_type"] = c.DictType
	}

	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// GetDBColumnList admin-从数据库表中获取表字段列表
func (e *SysGenColumn) GetDBColumnList(dbTableName string) ([]models.DBColumn, int, error) {
	if len(dbTableName) <= 0 {
		return nil, baseLang.SysGenTableSelectCode, lang.MsgErr(baseLang.SysGenTableSelectCode, e.Lang)
	}
	var data []models.DBColumn
	err := e.Orm.Where("TABLE_NAME in (?)", dbTableName).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Delete admin-删除表字段
func (e *SysGenColumn) Delete(req dto.SysGenColumnDeleteReq, p *middleware.DataPermission) (int, error) {
	var err error
	if len(req.Ids) > 0 {
		var data models.SysGenColumn
		err = e.Orm.Scopes(
			middleware.Permission(data.TableName(), p),
		).Delete(&data, req.Ids).Error
		if err != nil {
			return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
		}
	}
	if len(req.TableIds) > 0 {
		var data models.SysGenColumn
		err = e.Orm.Scopes(
			middleware.Permission(data.TableName(), p),
		).Where("table_id in (?)", req.TableIds).Delete(&models.SysGenColumn{}).Error
		if err != nil {
			return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
		}
	}
	return baseLang.SuccessCode, nil
}
