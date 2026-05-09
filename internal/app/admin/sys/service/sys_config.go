package service

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"go-admin/internal/app/admin/sys/models"
	"go-admin/internal/app/admin/sys/service/dto"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/dateutils"
	"gorm.io/gorm"
	"strconv"
	"time"

	cDto "go-admin/pkg/dto"
)

type SysConfig struct {
	service.Service
}

// NewSysConfigService admin-实例化配置管理
func NewSysConfigService(s *service.Service) *SysConfig {
	var srv = new(SysConfig)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取配置管理分页列表
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
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// GetList admin-获取系统配置全部列表
func (e *SysConfig) GetList(c *dto.SysConfigQueryReq) ([]models.SysConfig, int, error) {
	var list []models.SysConfig
	var err error
	err = e.Orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
	).Find(&list).Error
	if err != nil {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, clang.SuccessCode, nil
}

// Get admin-获取配置管理详情
func (e *SysConfig) Get(id int64, p *middleware.DataPermission) (*models.SysConfig, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.SysConfig{}
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

// QueryOne admin-获取配置管理一条记录
func (e *SysConfig) QueryOne(queryCondition *dto.SysConfigQueryReq, p *middleware.DataPermission) (*models.SysConfig, int, error) {
	data := &models.SysConfig{}
	err := e.Orm.Model(&models.SysConfig{}).
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

// Count admin-获取配置管理数据总数
func (e *SysConfig) Count(c *dto.SysConfigQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysConfig{}).
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

// Insert admin-新增配置管理
func (e *SysConfig) Insert(c *dto.SysConfigInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.ConfigName == "" {
		return 0, clang.SysConfNameEmptyCode, lang.MsgErr(clang.SysConfNameEmptyCode, e.Lang)
	}
	if c.ConfigKey == "" {
		return 0, clang.SysConfKeyEmptyCode, lang.MsgErr(clang.SysConfKeyEmptyCode, e.Lang)
	}
	if c.ConfigValue == "" {
		return 0, clang.SysConfValueEmptyCode, lang.MsgErr(clang.SysConfValueEmptyCode, e.Lang)
	}
	if c.ConfigType == "" {
		return 0, clang.SysConfTypeEmptyCode, lang.MsgErr(clang.SysConfTypeEmptyCode, e.Lang)
	}
	if c.IsFrontend == "" {
		return 0, clang.SysConfIsFrontendEmptyCode, lang.MsgErr(clang.SysConfIsFrontendEmptyCode, e.Lang)
	}

	req := dto.SysConfigQueryReq{}
	req.ConfigKey = c.ConfigKey
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != clang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, clang.SysConfKeyExistCode, lang.MsgErr(clang.SysConfKeyExistCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update admin-更新配置管理
func (e *SysConfig) Update(c *dto.SysConfigUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.ConfigName == "" {
		return false, clang.SysConfNameEmptyCode, lang.MsgErr(clang.SysConfNameEmptyCode, e.Lang)
	}
	if c.ConfigKey == "" {
		return false, clang.SysConfKeyEmptyCode, lang.MsgErr(clang.SysConfKeyEmptyCode, e.Lang)
	}
	if c.ConfigValue == "" {
		return false, clang.SysConfValueEmptyCode, lang.MsgErr(clang.SysConfValueEmptyCode, e.Lang)
	}
	if c.ConfigType == "" {
		return false, clang.SysConfTypeEmptyCode, lang.MsgErr(clang.SysConfTypeEmptyCode, e.Lang)
	}
	if c.IsFrontend == "" {
		return false, clang.SysConfIsFrontendEmptyCode, lang.MsgErr(clang.SysConfIsFrontendEmptyCode, e.Lang)
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
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysDictDataValueExistCode, lang.MsgErr(clang.SysDictDataValueExistCode, e.Lang)
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
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete admin-删除配置管理
func (e *SysConfig) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysConfig
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
}

// Export admin-导出配置管理
func (e *SysConfig) Export(list []models.SysConfig) ([]byte, error) {
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
		configType := dictService.GetLabel("admin_sys_config_type", item.ConfigType)
		isFrontend := dictService.GetLabel("admin_sys_config_is_frontend", item.IsFrontend)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.ConfigName, item.ConfigKey, item.ConfigValue, item.Remark, configType, isFrontend, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// GetByKey admin-根据Key获取配置
func (e *SysConfig) GetByKey(c *dto.SysConfigByKeyReq) (*dto.SysConfigByKeyResp, int, error) {
	var err error
	var data models.SysConfig
	resp := &dto.SysConfigByKeyResp{}
	err = e.Orm.Scopes().Table(data.TableName()).Where("config_key = ?", c.ConfigKey).First(resp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return resp, clang.SuccessCode, nil
}

// GetWithKeyStr admin-使用字符串key获取配置
func (e *SysConfig) GetWithKeyStr(key string) (string, int, error) {
	query := dto.SysConfigByKeyReq{}
	query.ConfigKey = key

	resp, respCode, err := e.GetByKey(&query)
	if err != nil || resp.ConfigValue == "" {
		return "", respCode, err
	}
	return resp.ConfigValue, clang.SuccessCode, nil
}

// GetWithKeyInt admin-使用数字key获取配置
func (e *SysConfig) GetWithKeyInt(key string) (int, int, error) {
	query := dto.SysConfigByKeyReq{}
	query.ConfigKey = key

	resp, respCode, err := e.GetByKey(&query)
	if err != nil || resp.ConfigValue == "" {
		return -1, respCode, err
	}
	value, err := strconv.ParseInt(resp.ConfigValue, 10, 64)
	if err != nil {
		return -1, clang.SysConfGetErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.SysConfGetErrCode, clang.SysConfGetErrLogCode, err)
	}
	return int(value), clang.SuccessCode, nil
}

// GetWithKeyDecimal admin-使用字符串key获取配置，返回decimal
func (e *SysConfig) GetWithKeyDecimal(key string) (*decimal.Decimal, int, error) {
	resultValue, respCode, err := e.GetWithKeyStr(key)
	if err != nil {
		return nil, respCode, err
	}
	result, err := decimal.NewFromString(resultValue)
	if err != nil {
		return nil, clang.SysConfGetErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.SysConfGetErrCode, clang.SysConfGetErrLogCode, err)
	}
	return &result, clang.SuccessCode, nil

}
