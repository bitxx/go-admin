package service

import (
	sysLang "go-admin/app/admin/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/runtime"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/core/dto"
)

type SysDictData struct {
	service.Service
}

func NewSysDictDataService(s *service.Service) *SysDictData {
	var srv = new(SysDictData)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取列表
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

func (e *SysDictData) QueryOne(queryCondition *dto.SysDictDataQueryReq, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	data := &models.SysDictData{}
	err := e.Orm.Model(&models.SysDictData{}).
		Scopes(
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

func (e *SysDictData) Count(c *dto.SysDictDataQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Get 获取对象
func (e *SysDictData) Get(id int64, p *middleware.DataPermission) (*models.SysDictData, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysDictData{}
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

// Insert 创建对象
func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.DictLabel == "" {
		return 0, sysLang.SysDictDataLabelEmptyCode, lang.MsgErr(sysLang.SysDictDataLabelEmptyCode, e.Lang)
	}
	if c.DictValue == "" {
		return 0, sysLang.SysDictDataValueEmptyCode, lang.MsgErr(sysLang.SysDictDataValueEmptyCode, e.Lang)
	}
	if c.DictSort < 0 {
		return 0, sysLang.SysDictDataSortEmptyCode, lang.MsgErr(sysLang.SysDictDataSortEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return 0, sysLang.SysDictTypeTypeEmptyCode, lang.MsgErr(sysLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	req := dto.SysDictDataQueryReq{}
	req.DictType = c.DictType
	req.DictValue = c.DictValue
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysDictDataValueExistCode, lang.MsgErr(sysLang.SysDictDataValueExistCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改对象
func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
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
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDictDataValueExistCode, lang.MsgErr(sysLang.SysDictDataValueExistCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateDictType 用于sys_dict_type同步修改
func (e *SysDictData) UpdateDictType(oldDictType, newDictType string) (int, error) {
	if oldDictType == newDictType {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	updates := map[string]interface{}{}
	updates["dict_type"] = newDictType
	err := e.Orm.Model(&models.SysDictData{}).Where("dict_type=?", oldDictType).Updates(&updates).Error
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	return lang.SuccessCode, nil
}

// Remove 删除
func (e *SysDictData) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	var err error
	var data models.SysDictData
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetAll 获取所有
func (e *SysDictData) GetAll(c *dto.SysDictDataQueryReq) ([]models.SysDictData, int, error) {
	var err error
	var list []models.SysDictData

	err = e.Orm.Model(&models.SysDictData{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

// GetValue 根据dict和key获取值
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
