package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	sysLang "go-admin/app/admin/lang"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/global"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateUtils"
	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type SysDictType struct {
	service.Service
}

func NewSysDictTypeService(s *service.Service) *SysDictType {
	var srv = new(SysDictType)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取列表
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

func (e *SysDictType) QueryOne(queryCondition *dto.SysDictTypeQueryReq, p *middleware.DataPermission) (*models.SysDictType, int, error) {
	data := &models.SysDictType{}
	err := e.Orm.Model(&models.SysDictType{}).
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

func (e *SysDictType) Count(c *dto.SysDictTypeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDictType{}).
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
func (e *SysDictType) Get(id int64, p *middleware.DataPermission) (*models.SysDictType, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysDictType{}
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
func (e *SysDictType) Insert(c *dto.SysDictTypeInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.DictName == "" {
		return 0, sysLang.SysDictTypeNameEmptyCode, lang.MsgErr(sysLang.SysDictTypeNameEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return 0, sysLang.SysDictTypeTypeEmptyCode, lang.MsgErr(sysLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	req := dto.SysDictTypeQueryReq{}
	req.DictType = c.DictType
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysDictTypeTypeExistCode, lang.MsgErr(sysLang.SysDictTypeTypeExistCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改对象
func (e *SysDictType) Update(c *dto.SysDictTypeUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.DictName == "" {
		return false, sysLang.SysDictTypeNameEmptyCode, lang.MsgErr(sysLang.SysDictTypeNameEmptyCode, e.Lang)
	}
	if c.DictType == "" {
		return false, sysLang.SysDictTypeTypeEmptyCode, lang.MsgErr(sysLang.SysDictTypeTypeEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.DictName != "" && data.DictName != c.DictName {
		updates["dict_name"] = c.DictName
	}
	//不得修改类型
	if c.DictType != "" && data.DictType != c.DictType {
		//判断是否已存在
		req := dto.SysDictTypeQueryReq{}
		req.DictType = c.DictType
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDictTypeTypeExistCode, lang.MsgErr(sysLang.SysDictTypeTypeExistCode, e.Lang)
		}
		updates["dict_type"] = c.DictType
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

// Remove 删除
func (e *SysDictType) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
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
		if err != nil && respCode != lang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return sysLang.SysDictTypeTypeHasUsedCode, lang.MsgErr(sysLang.SysDictTypeTypeHasUsedCode, e.Lang)
		}
	}

	var err error
	var data models.SysDictType
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetAll 获取所有
func (e *SysDictType) GetAll(c *dto.SysDictTypeQueryReq) ([]models.SysDictType, int, error) {
	var err error
	var data models.SysDictType
	var list []models.SysDictType

	err = e.Orm.Model(&data).Order("created_at desc").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

// GetExcel 导出Category
func (e *SysDictType) GetExcel(list []models.SysDictType) ([]byte, error) {
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
			item.Id, item.DictName, item.DictType, item.Remark, dateUtils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
