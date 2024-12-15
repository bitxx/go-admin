package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"

	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"

	"gorm.io/gorm"
	"time"
)

type ContentArticle struct {
	service.Service
}

// NewContentArticleService plugins-实例化文章管理
func NewContentArticleService(s *service.Service) *ContentArticle {
	var srv = new(ContentArticle)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage plugins-获取文章管理分页列表
func (e *ContentArticle) GetPage(c *dto.ContentArticleQueryReq, p *middleware.DataPermission) ([]models.ContentArticle, int64, int, error) {
	var data models.ContentArticle
	var list []models.ContentArticle
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).Preload("ContentCategory").
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

// Get plugins-获取文章管理详情
func (e *ContentArticle) Get(id int64, p *middleware.DataPermission) (*models.ContentArticle, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.ContentArticle{}
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

// QueryOne plugins-获取文章管理一条记录
func (e *ContentArticle) QueryOne(queryCondition *dto.ContentArticleQueryReq, p *middleware.DataPermission) (*models.ContentArticle, int, error) {
	data := &models.ContentArticle{}
	err := e.Orm.Scopes(
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

// Count admin-获取文章管理数据总数
func (e *ContentArticle) Count(queryCondition *dto.ContentArticleQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentArticle{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert plugins-新增文章管理
func (e *ContentArticle) Insert(c *dto.ContentArticleInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return 0, baseLang.PluginsArticleNameCode, lang.MsgErr(baseLang.PluginsArticleNameCode, e.Lang)
	}
	if c.Content == "" {
		return 0, baseLang.PluginsArticleContentCode, lang.MsgErr(baseLang.PluginsArticleContentCode, e.Lang)
	}
	if c.CateId <= 0 {
		return 0, baseLang.PluginsArticleCatIdEmptyCode, lang.MsgErr(baseLang.PluginsArticleCatIdEmptyCode, e.Lang)
	}

	//确保文章名称不存在
	req := dto.ContentArticleQueryReq{}
	req.NameInner = c.Name
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.PluginsArticleNameHasUsedCode, lang.MsgErr(baseLang.PluginsArticleNameHasUsedCode, e.Lang)
	}

	//确保分类存在
	categoryService := NewContentCategoryService(&e.Service)
	_, respCode, err = categoryService.Get(c.CateId, nil)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if err != nil && respCode == baseLang.DataNotFoundCode {
		return 0, baseLang.PluginsCategoryNotFoundCode, lang.MsgErr(baseLang.PluginsCategoryNotFoundCode, e.Lang)
	}
	now := time.Now()
	var data models.ContentArticle
	data.CateId = c.CateId
	data.Name = c.Name
	data.Content = c.Content
	data.Remark = c.Remark
	data.Status = c.Status
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

// Update plugins-更新文章管理
func (e *ContentArticle) Update(c *dto.ContentArticleUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return false, baseLang.PluginsArticleNameCode, lang.MsgErr(baseLang.PluginsArticleNameCode, e.Lang)
	}
	if c.Content == "" {
		return false, baseLang.PluginsArticleContentCode, lang.MsgErr(baseLang.PluginsArticleContentCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.CateId > 0 && data.CateId != c.CateId {
		updates["cate_id"] = c.CateId
	}
	if c.Name != "" && data.Name != c.Name {
		req := dto.ContentArticleQueryReq{}
		req.NameInner = c.Name
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysDictDataValueExistCode, lang.MsgErr(baseLang.SysDictDataValueExistCode, e.Lang)
		}
		updates["name"] = c.Name
	}
	if c.Content != "" && data.Content != c.Content {
		updates["content"] = c.Content
	}
	if c.Remark != "" && data.Remark != c.Remark {
		updates["remark"] = c.Remark
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
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

// Delete plugins-删除文章管理
func (e *ContentArticle) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.ContentArticle
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export plugins-导出文章管理
func (e *ContentArticle) Export(list []models.ContentArticle) ([]byte, error) {
	sheetName := "ContentArticle"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "P", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"文章编号", "分类名称", "标题", "内容", "备注", "时间"})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.ContentCategory.Name, item.Name, item.Content, item.Remark, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
