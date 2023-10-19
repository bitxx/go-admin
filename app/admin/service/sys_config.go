package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	sysLang "go-admin/app/admin/lang"
	"go-admin/common/dto/service"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"gorm.io/gorm"
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type SysConfig struct {
	service.Service
}

func NewSysConfigService(s *service.Service) *SysConfig {
	var srv = new(SysConfig)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(c *dto.SysConfigQueryReq, p *middleware.DataPermission) ([]models.SysConfig, int64, int, error) {
	var list []models.SysConfig
	var data models.SysConfig
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

// Get 获取SysConfig对象
func (e *SysConfig) Get(id int64, p *middleware.DataPermission) (*models.SysConfig, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysConfig{}
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

func (e *SysConfig) QueryOne(queryCondition *dto.SysConfigQueryReq, p *middleware.DataPermission) (*models.SysConfig, int, error) {
	data := &models.SysConfig{}
	err := e.Orm.Model(&models.SysConfig{}).
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

// Count 获取条数
func (e *SysConfig) Count(c *dto.SysConfigQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysConfig{}).
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

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *dto.SysConfigInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.ConfigName == "" {
		return 0, sysLang.SysConfNameEmptyCode, lang.MsgErr(sysLang.SysConfNameEmptyCode, e.Lang)
	}
	if c.ConfigKey == "" {
		return 0, sysLang.SysConfKeyEmptyCode, lang.MsgErr(sysLang.SysConfKeyEmptyCode, e.Lang)
	}
	if c.ConfigValue == "" {
		return 0, sysLang.SysConfValueEmptyCode, lang.MsgErr(sysLang.SysConfValueEmptyCode, e.Lang)
	}
	if c.ConfigType == "" {
		return 0, sysLang.SysConfTypeEmptyCode, lang.MsgErr(sysLang.SysConfTypeEmptyCode, e.Lang)
	}
	if c.IsFrontend == "" {
		return 0, sysLang.SysConfIsFrontendEmptyCode, lang.MsgErr(sysLang.SysConfIsFrontendEmptyCode, e.Lang)
	}

	req := dto.SysConfigQueryReq{}
	req.ConfigKey = c.ConfigKey
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysConfKeyExistCode, lang.MsgErr(sysLang.SysConfKeyExistCode, e.Lang)
	}

	now := time.Now()
	data := models.SysConfig{}
	data.ConfigName = c.ConfigName
	data.ConfigKey = c.ConfigKey
	data.ConfigValue = c.ConfigValue
	data.ConfigType = c.ConfigType
	data.IsFrontend = c.IsFrontend
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

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *dto.SysConfigUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.ConfigName == "" {
		return false, sysLang.SysConfNameEmptyCode, lang.MsgErr(sysLang.SysConfNameEmptyCode, e.Lang)
	}
	if c.ConfigKey == "" {
		return false, sysLang.SysConfKeyEmptyCode, lang.MsgErr(sysLang.SysConfKeyEmptyCode, e.Lang)
	}
	if c.ConfigValue == "" {
		return false, sysLang.SysConfValueEmptyCode, lang.MsgErr(sysLang.SysConfValueEmptyCode, e.Lang)
	}
	if c.ConfigType == "" {
		return false, sysLang.SysConfTypeEmptyCode, lang.MsgErr(sysLang.SysConfTypeEmptyCode, e.Lang)
	}
	if c.IsFrontend == "" {
		return false, sysLang.SysConfIsFrontendEmptyCode, lang.MsgErr(sysLang.SysConfIsFrontendEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}

	if c.ConfigName != "" && data.ConfigName != c.ConfigName {
		updates["config_name"] = c.ConfigName
	}
	if c.ConfigKey != "" && data.ConfigKey != c.ConfigKey {
		//检测是否重复
		req := dto.SysConfigQueryReq{}
		req.ConfigKey = c.ConfigKey
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDictDataValueExistCode, lang.MsgErr(sysLang.SysDictDataValueExistCode, e.Lang)
		}
		updates["config_key"] = c.ConfigKey
	}
	if c.ConfigValue != "" && data.ConfigValue != c.ConfigValue {
		updates["config_value"] = c.ConfigValue
	}
	if c.ConfigType != "" && data.ConfigType != c.ConfigType {
		updates["config_type"] = c.ConfigType
	}
	if c.IsFrontend != "" && data.IsFrontend != c.IsFrontend {
		updates["is_frontend"] = c.IsFrontend
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

// Remove 删除SysConfig
func (e *SysConfig) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysConfig
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel 导出配置
func (e *SysConfig) GetExcel(list []models.SysConfig) ([]byte, error) {
	//sheet名称
	sheetName := "config"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "H", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"配置编号", "配置名称", "键名", "键值", "配置类型", "是否前端展示", "备注", "创建时间"})
	dictService := NewSysDictDataService(&e.Service)

	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		configType := dictService.GetLabel("sys_config_type", item.ConfigType)
		isFrontend := dictService.GetLabel("sys_config_is_frontend", item.IsFrontend)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.ConfigName, item.ConfigKey, item.ConfigValue, item.Remark, configType, isFrontend, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// GetByKey 根据Key获取SysConfig
func (e *SysConfig) GetByKey(c *dto.SysConfigByKeyReq) (*dto.SysConfigByKeyResp, int, error) {
	var err error
	var data models.SysConfig
	resp := &dto.SysConfigByKeyResp{}
	err = e.Orm.Scopes().Table(data.TableName()).Where("config_key = ?", c.ConfigKey).First(resp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return resp, lang.SuccessCode, nil
}

// GetWithKeyStr 使用字符串类型快速获取配置结果
func (e *SysConfig) GetWithKeyStr(key string) (string, int, error) {
	query := dto.SysConfigByKeyReq{}
	query.ConfigKey = key

	resp, respCode, err := e.GetByKey(&query)
	if err != nil || resp.ConfigValue == "" {
		return "", respCode, err
	}
	return resp.ConfigValue, lang.SuccessCode, nil
}

func (e *SysConfig) GetWithKeyList(c *dto.SysConfigQueryReq) ([]models.SysConfig, int, error) {
	var list []models.SysConfig
	var err error
	err = e.Orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
	).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

func (e *SysConfig) GetWithKeyInt(key string) (int, int, error) {
	query := dto.SysConfigByKeyReq{}
	query.ConfigKey = key

	resp, respCode, err := e.GetByKey(&query)
	if err != nil || resp.ConfigValue == "" {
		return -1, respCode, err
	}
	value, err := strconv.ParseInt(resp.ConfigValue, 10, 64)
	if err != nil {
		return -1, sysLang.SysConfGetErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, sysLang.SysConfGetErrCode, sysLang.SysConfGetErrLogCode, err)
	}
	return int(value), lang.SuccessCode, nil
}

// GetWithKeyDecimal 返回
func (e *SysConfig) GetWithKeyDecimal(key string) (*decimal.Decimal, int, error) {
	resultValue, respCode, err := e.GetWithKeyStr(key)
	if err != nil {
		return nil, respCode, err
	}
	result, err := decimal.NewFromString(resultValue)
	if err != nil {
		return nil, sysLang.SysConfGetErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, sysLang.SysConfGetErrCode, sysLang.SysConfGetErrLogCode, err)
	}
	return &result, lang.SuccessCode, nil

}
