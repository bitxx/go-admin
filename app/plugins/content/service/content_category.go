package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"

	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
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

type ContentCategory struct {
	service.Service
}

// NewContentCategoryService plugins-实例化内容分类管理
func NewContentCategoryService(s *service.Service) *ContentCategory {
	var srv = new(ContentCategory)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage plugins-获取内容分类管理分页列表
func (e *ContentCategory) GetPage(c *dto.ContentCategoryQueryReq, p *middleware.DataPermission) ([]models.ContentCategory, int64, int, error) {
	var data models.ContentCategory
	var list []models.ContentCategory
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

// Get plugins-获取内容分类管理详情
func (e *ContentCategory) Get(id int64, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.ContentCategory{}
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

// QueryOne plugins-获取内容分类管理一条记录
func (e *ContentCategory) QueryOne(queryCondition *dto.ContentCategoryQueryReq, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	data := &models.ContentCategory{}
	err := e.Orm.Scopes(
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

// Count admin-获取内容分类管理数据总数
func (e *ContentCategory) Count(queryCondition *dto.ContentCategoryQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentCategory{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert plugins-新增内容分类管理详情
func (e *ContentCategory) Insert(c *dto.ContentCategoryInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return 0, baseLang.PluginsCategoryNameCode, lang.MsgErr(baseLang.PluginsCategoryNameCode, e.Lang)
	}
	req := dto.ContentCategoryQueryReq{}
	req.NameInner = c.Name
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.PluginsCategoryNameHasUsedCode, lang.MsgErr(baseLang.PluginsCategoryNameHasUsedCode, e.Lang)
	}
	now := time.Now()
	var data models.ContentCategory
	data.Name = c.Name
	data.Status = global.SysStatusOk
	data.Remark = c.Remark
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	data.CreatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update plugins-更新内容分类管理
func (e *ContentCategory) Update(c *dto.ContentCategoryUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return false, baseLang.PluginsCategoryNameCode, lang.MsgErr(baseLang.PluginsCategoryNameCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}

	if c.Name != "" && data.Name != c.Name {
		req := dto.ContentCategoryQueryReq{}
		req.NameInner = c.Name
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.PluginsCategoryNameHasUsedCode, lang.MsgErr(baseLang.PluginsCategoryNameHasUsedCode, e.Lang)
		}
		updates["name"] = c.Name
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

// Delete plugins-删除内容分类管理
func (e *ContentCategory) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	//若有文章，不得删除
	articleService := NewContentArticleService(&e.Service)
	articleReq := dto.ContentArticleQueryReq{}
	articleReq.CateIds = ids
	count, respCode, err := articleService.Count(&articleReq)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return baseLang.PluginsCategoryNameHasUsedCode, lang.MsgErr(baseLang.PluginsCategoryNameHasUsedCode, e.Lang)
	}
	var data models.ContentCategory
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export plugins-导出内容分类管理
func (e *ContentCategory) Export(list []models.ContentCategory) ([]byte, error) {
	sheetName := "ContentCategory"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "P", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"分类编号", "分类名称", "创建时间"})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.Name, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
