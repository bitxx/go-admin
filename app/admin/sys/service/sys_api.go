package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"

	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/runtime"
	"go-admin/core/utils/dateutils"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysApi struct {
	service.Service
}

// NewSysApiService admin-实例化接口管理
func NewSysApiService(s *service.Service) *SysApi {
	var srv = new(SysApi)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取接口管理分页列表
func (e *SysApi) GetPage(c *dto.SysApiQueryReq, p *middleware.DataPermission) ([]models.SysApi, int64, int, error) {
	var list []models.SysApi
	var data models.SysApi
	var count int64

	err := e.Orm.Order("id desc").Model(&data).Preload("SysMenu").
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

// GetList admin-获取接口管理全部列表
func (e *SysApi) GetList(c *dto.SysApiQueryReq, p *middleware.DataPermission) ([]models.SysApi, int64, int, error) {
	var list []models.SysApi
	var data models.SysApi
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get admin-获取接口管理详情
func (e *SysApi) Get(id int64, p *middleware.DataPermission) (*models.SysApi, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysApi{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne admin-获取接口管理一条记录
func (e *SysApi) QueryOne(queryCondition *dto.SysApiQueryReq, p *middleware.DataPermission) (*models.SysApi, int, error) {
	data := &models.SysApi{}
	err := e.Orm.Model(&models.SysApi{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Update admin-更新接口管理
func (e *SysApi) Update(c *dto.SysApiUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.ApiType != "" && data.ApiType != c.ApiType {
		updates["api_type"] = c.ApiType
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

// Delete admin-删除接口管理
func (e *SysApi) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysApi

	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		// 删除子表数据
		if err := tx.Table("admin_sys_menu_api_rule").Where("admin_sys_api_id in (?)", ids).Delete(nil).Error; err != nil {
			return err
		}
		err = tx.Scopes(
			middleware.Permission(data.TableName(), p),
		).Delete(&data, ids).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export admin-导出接口管理
func (e *SysApi) Export(list []models.SysApi) ([]byte, error) {
	//sheet名称
	sheetName := "Api"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "F", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "标题", "请求地址", "请求方法", "请求类型", "创建时间"})

	dictService := NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		method := dictService.GetLabel("admin_sys_api_method", item.Method)
		apiType := dictService.GetLabel("admin_sys_config_type", item.ApiType)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.Description, item.Path, method, apiType, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// Sync admin-接口同步数据
func (e *SysApi) Sync() (int, error) {
	var routers = runtime.RuntimeConfig.GetRouter()

	err := models.SaveSysApi(e.Orm, routers)
	if err != nil {
		return baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
