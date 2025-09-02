package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"gorm.io/gorm"
	"time"
)

type SysDictType struct {
	service.Service
}

// NewSysDictTypeService admin-实例化字典类型
func NewSysDictTypeService(s *service.Service) *SysDictType {
	var srv = new(SysDictType)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取字典类型分页列表
func (e *SysDictType) GetPage(c *dto.SysDictTypeQueryReq, p *middleware.DataPermission) ([]models.SysDictType, int64, int, error) {
	var list []models.SysDictType
	var data models.SysDictType
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

// QueryOne admin-获取字典类型一条记录
func (e *SysDictType) QueryOne(queryCondition *dto.SysDictTypeQueryReq, p *middleware.DataPermission) (*models.SysDictType, int, error) {
	data := &models.SysDictType{}
	err := e.Orm.Model(&models.SysDictType{}).
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

// Count admin-获取字典类型数据总数
func (e *SysDictType) Count(c *dto.SysDictTypeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDictType{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Get admin-获取字典类型详情
func (e *SysDictType) Get(id int64, p *middleware.DataPermission) (*models.SysDictType, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysDictType{}
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

// Insert admin-新增字典类型
func (e *SysDictType) Insert(c *dto.SysDictTypeInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.DictName == "" {
		return 0, baseLang.SysDictTypeNameEmptyCode, lang.MsgErr(baseLang.SysDictTypeNameEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return 0, baseLang.SysDictTypeTypeEmptyCode, lang.MsgErr(baseLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	req := dto.SysDictTypeQueryReq{}
	req.DictType = c.DictType
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.SysDictTypeTypeExistCode, lang.MsgErr(baseLang.SysDictTypeTypeExistCode, e.Lang)
	}

	now := time.Now()
	data := models.SysDictType{}
	data.DictName = c.DictName
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

// Update admin-更新字典类型
func (e *SysDictType) Update(c *dto.SysDictTypeUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.DictName == "" {
		return false, baseLang.SysDictTypeNameEmptyCode, lang.MsgErr(baseLang.SysDictTypeNameEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return false, baseLang.SysDictTypeTypeEmptyCode, lang.MsgErr(baseLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	updates := map[string]interface{}{}
	if c.DictName != "" && data.DictName != c.DictName {
		updates["dict_name"] = c.DictName
	}
	//不得修改类型
	if c.DictType != "" && data.DictType != c.DictType {
		//判断是否已存在
		req := dto.SysDictTypeQueryReq{}
		req.DictType = c.DictType
		resp, respCode, er := e.QueryOne(&req, nil)
		if er != nil && respCode != baseLang.DataNotFoundCode {
			err = er
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysDictTypeTypeExistCode, lang.MsgErr(baseLang.SysDictTypeTypeExistCode, e.Lang)
		}
		updates["dict_type"] = c.DictType
		dictDataService := NewSysDictDataService(&e.Service)
		respCode, err = dictDataService.UpdateDictType(data.DictType, c.DictType)
		if err != nil {
			return false, respCode, err
		}
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

// Delete admin-删除字典类型
func (e *SysDictType) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	for _, id := range ids {
		dictType, reqCode, err := e.Get(id, p)
		if err != nil {
			return reqCode, err
		}
		//若被使用，不得删除
		dataService := NewSysDictDataService(&e.Service)
		dataReq := dto.SysDictDataQueryReq{}
		dataReq.DictType = dictType.DictType
		count, respCode, err := dataService.Count(&dataReq)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return baseLang.SysDictTypeTypeHasUsedCode, lang.MsgErr(baseLang.SysDictTypeTypeHasUsedCode, e.Lang)
		}
	}

	var err error
	var data models.SysDictType
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// GetList admin-获取字典类型全部列表
func (e *SysDictType) GetList(c *dto.SysDictTypeQueryReq) ([]models.SysDictType, int, error) {
	var err error
	var data models.SysDictType
	var list []models.SysDictType

	err = e.Orm.Model(&data).Order("created_at desc").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, baseLang.SuccessCode, nil
}

// Export admin-导出字典类型
func (e *SysDictType) Export(list []models.SysDictType) ([]byte, error) {
	//sheet名称
	sheetName := "DictType"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "E", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"字典编号", "字典名称", "字典类型", "备注", "创建时间"})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.DictName, item.DictType, item.Remark, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
