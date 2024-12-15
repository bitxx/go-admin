package service

import (
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/runtime"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
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
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// QueryOne admin-获取字典数据一条记录
func (e *SysDictData) QueryOne(queryCondition *dto.SysDictDataQueryReq, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	data := &models.SysDictData{}
	err := e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
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
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Get admin-获取字典数据详情
func (e *SysDictData) Get(id int64, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysDictData{}
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

// Insert admin-新增字典数据
func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.DictLabel == "" {
		return 0, baseLang.SysDictDataLabelEmptyCode, lang.MsgErr(baseLang.SysDictDataLabelEmptyCode, e.Lang)
	}
	if c.DictValue == "" {
		return 0, baseLang.SysDictDataValueEmptyCode, lang.MsgErr(baseLang.SysDictDataValueEmptyCode, e.Lang)
	}
	if c.DictSort < 0 {
		return 0, baseLang.SysDictDataSortEmptyCode, lang.MsgErr(baseLang.SysDictDataSortEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return 0, baseLang.SysDictTypeTypeEmptyCode, lang.MsgErr(baseLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	req := dto.SysDictDataQueryReq{}
	req.DictType = c.DictType
	req.DictValue = c.DictValue
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.SysDictDataValueExistCode, lang.MsgErr(baseLang.SysDictDataValueExistCode, e.Lang)
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
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update admin-更新字典数据
func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysDictDataValueExistCode, lang.MsgErr(baseLang.SysDictDataValueExistCode, e.Lang)
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
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// UpdateDictType 用于admin_sys_dict_type同步修改
func (e *SysDictData) UpdateDictType(oldDictType, newDictType string) (int, error) {
	if oldDictType == newDictType {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	updates := map[string]interface{}{}
	updates["dict_type"] = newDictType
	err := e.Orm.Model(&models.SysDictData{}).Where("dict_type=?", oldDictType).Updates(&updates).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Delete admin-删除字典数据
func (e *SysDictData) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	var err error
	var data models.SysDictData
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
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
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, baseLang.SuccessCode, nil
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
