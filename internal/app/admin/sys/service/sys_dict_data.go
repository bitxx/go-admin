package service

import (
	"errors"
	"go-admin/internal/app/admin/sys/models"
	"go-admin/internal/app/admin/sys/service/dto"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/runtime"
	"gorm.io/gorm"
	"time"

	cDto "go-admin/pkg/dto"
)

type SysDictData struct {
	service.Service
}

// NewSysDictDataService admin-实例化字典数据
func NewSysDictDataService(s *service.Service) *SysDictData {
	var srv = new(SysDictData)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取字典数据分页列表
func (e *SysDictData) GetPage(c *dto.SysDictDataQueryReq, p *middleware.DataPermission) ([]models.SysDictData, int64, int, error) {
	var list []models.SysDictData
	var data models.SysDictData
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// QueryOne admin-获取字典数据一条记录
func (e *SysDictData) QueryOne(queryCondition *dto.SysDictDataQueryReq, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	data := &models.SysDictData{}
	err := e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// Count admin-获取字典数据数据总数
func (e *SysDictData) Count(c *dto.SysDictDataQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return count, clang.SuccessCode, nil
}

// Get admin-获取字典数据详情
func (e *SysDictData) Get(id int64, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.SysDictData{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// Insert admin-新增字典数据
func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.DictLabel == "" {
		return 0, clang.SysDictDataLabelEmptyCode, lang.MsgErr(clang.SysDictDataLabelEmptyCode, e.Lang)
	}
	if c.DictValue == "" {
		return 0, clang.SysDictDataValueEmptyCode, lang.MsgErr(clang.SysDictDataValueEmptyCode, e.Lang)
	}
	if c.DictSort < 0 {
		return 0, clang.SysDictDataSortEmptyCode, lang.MsgErr(clang.SysDictDataSortEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return 0, clang.SysDictTypeTypeEmptyCode, lang.MsgErr(clang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	req := dto.SysDictDataQueryReq{}
	req.DictType = c.DictType
	req.DictValue = c.DictValue
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != clang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, clang.SysDictDataValueExistCode, lang.MsgErr(clang.SysDictDataValueExistCode, e.Lang)
	}

	now := time.Now()
	data := models.SysDictData{}
	data.DictSort = c.DictSort
	data.DictLabel = c.DictLabel
	data.DictValue = c.DictValue
	data.DictType = c.DictType
	data.Status = global.SysStatusOk
	data.Remark = c.Remark
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update admin-更新字典数据
func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}

	if c.DictSort.IntPart() > 0 && data.DictSort != int(c.DictSort.IntPart()) {
		updates["dict_sort"] = c.DictSort
	}
	if c.DictLabel != "" && data.DictLabel != c.DictLabel {
		updates["dict_label"] = c.DictLabel
	}
	if c.DictValue != "" && data.DictValue != c.DictValue {
		req := dto.SysDictDataQueryReq{}
		req.DictType = c.DictType
		req.DictValue = c.DictValue
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysDictDataValueExistCode, lang.MsgErr(clang.SysDictDataValueExistCode, e.Lang)
		}
		updates["dict_value"] = c.DictValue
	}
	if c.Remark != "" && data.Remark != c.Remark {
		updates["remark"] = c.Remark
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// UpdateDictType 用于admin_sys_dict_type同步修改
func (e *SysDictData) UpdateDictType(oldDictType, newDictType string) (int, error) {
	if oldDictType == newDictType {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	updates := map[string]interface{}{}
	updates["dict_type"] = newDictType
	err := e.Orm.Model(&models.SysDictData{}).Where("dict_type=?", oldDictType).Updates(&updates).Error
	if err != nil {
		return clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
	}
	return clang.SuccessCode, nil
}

// Delete admin-删除字典数据
func (e *SysDictData) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	var err error
	var data models.SysDictData
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
}

// GetList admin-获取字典数据全部列表
func (e *SysDictData) GetList(c *dto.SysDictDataQueryReq) ([]models.SysDictData, int, error) {
	var err error
	var list []models.SysDictData

	err = e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, clang.SuccessCode, nil
}

// GetLabel admin-根据字典类型和键获取值
func (e *SysDictData) GetLabel(dict, value string) string {
	if dict == "" || value == "" {
		return ""
	}
	key := dict + value
	v, _ := runtime.RuntimeConfig.GetCacheAdapter().Get("", key)
	if v != "" {
		return v
	}

	var data models.SysDictData
	search := dto.SysDictDataQueryReq{}
	search.DictType = dict
	search.DictValue = value

	result := models.SysDictData{}

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(search.GetNeedSearch()),
		).First(&result).Error
	if err != nil {
		e.Log.Errorf("SysConfigService GetLabel error:%s", err)
		return ""
	}
	label := result.DictLabel
	//添加缓存
	_ = runtime.RuntimeConfig.GetCacheAdapter().Set("", key, label, -1)
	return label
}
