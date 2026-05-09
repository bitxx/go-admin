package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-admin/internal/app/plugins/content/models"
	"go-admin/internal/app/plugins/content/service/dto"
	clang "go-admin/internal/common/lang"
	cDto "go-admin/pkg/dto"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/dateutils"

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
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// Get plugins-获取内容分类管理详情
func (e *ContentCategory) Get(id int64, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.ContentCategory{}
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

// QueryOne plugins-获取内容分类管理一条记录
func (e *ContentCategory) QueryOne(queryCondition *dto.ContentCategoryQueryReq, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	data := &models.ContentCategory{}
	err := e.Orm.Scopes(
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

// Count admin-获取内容分类管理数据总数
func (e *ContentCategory) Count(queryCondition *dto.ContentCategoryQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentCategory{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return count, clang.SuccessCode, nil
}

// Insert plugins-新增内容分类管理详情
func (e *ContentCategory) Insert(c *dto.ContentCategoryInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return 0, clang.PluginsCategoryNameCode, lang.MsgErr(clang.PluginsCategoryNameCode, e.Lang)
	}
	req := dto.ContentCategoryQueryReq{}
	req.NameInner = c.Name
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != clang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, clang.PluginsCategoryNameHasUsedCode, lang.MsgErr(clang.PluginsCategoryNameHasUsedCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update plugins-更新内容分类管理
func (e *ContentCategory) Update(c *dto.ContentCategoryUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return false, clang.PluginsCategoryNameCode, lang.MsgErr(clang.PluginsCategoryNameCode, e.Lang)
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
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.PluginsCategoryNameHasUsedCode, lang.MsgErr(clang.PluginsCategoryNameHasUsedCode, e.Lang)
		}
		updates["name"] = c.Name
	}

	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete plugins-删除内容分类管理
func (e *ContentCategory) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	//若有文章，不得删除
	articleService := NewContentArticleService(&e.Service)
	articleReq := dto.ContentArticleQueryReq{}
	articleReq.CateIds = ids
	count, respCode, err := articleService.Count(&articleReq)
	if err != nil && respCode != clang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return clang.PluginsCategoryNameHasUsedCode, lang.MsgErr(clang.PluginsCategoryNameHasUsedCode, e.Lang)
	}
	var data models.ContentCategory
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
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
