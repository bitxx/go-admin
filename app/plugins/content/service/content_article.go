package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	sysLang "go-admin/app/admin/lang"
	cLang "go-admin/app/plugins/content/lang"
	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateUtils"

	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type ContentArticle struct {
	service.Service
}

// NewContentArticleService
// @Description: 实例化ContentArticle
// @param s
// @return *ContentArticle
func NewContentArticleService(s *service.Service) *ContentArticle {
	var srv = new(ContentArticle)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取ContentArticle列表
// @receiver e
// @param c
// @param p
// @return []models.ContentArticle
// @return int64
// @return int
// @return error
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get
// @Description: 获取ContentArticle对象
// @receiver e
// @param id 编号
// @param p
// @return *models.ContentArticle
// @return int
// @return error
func (e *ContentArticle) Get(id int64, p *middleware.DataPermission) (*models.ContentArticle, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.ContentArticle{}
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

// QueryOne
// @Description: 通过自定义条件获取ContentArticle一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.ContentArticle
// @return error
func (e *ContentArticle) QueryOne(queryCondition *dto.ContentArticleQueryReq, p *middleware.DataPermission) (*models.ContentArticle, int, error) {
	data := &models.ContentArticle{}
	err := e.Orm.Scopes(
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

// Count
//
//	@Description: 获取条数
//	@receiver e
//	@param c
//	@return int64
//	@return int
//	@return error
func (e *ContentArticle) Count(queryCondition *dto.ContentArticleQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentArticle{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Insert
// @Description: 创建ContentArticle对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *ContentArticle) Insert(c *dto.ContentArticleInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return 0, cLang.PluginsArticleNameCode, lang.MsgErr(cLang.PluginsArticleNameCode, e.Lang)
	}
	if c.Content == "" {
		return 0, cLang.PluginsArticleContentCode, lang.MsgErr(cLang.PluginsArticleContentCode, e.Lang)
	}
	if c.CateId <= 0 {
		return 0, cLang.PluginsArticleCatIdEmptyCode, lang.MsgErr(cLang.PluginsArticleCatIdEmptyCode, e.Lang)
	}

	//确保文章名称不存在
	req := dto.ContentArticleQueryReq{}
	req.NameInner = c.Name
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, cLang.PluginsArticleNameHasUsedCode, lang.MsgErr(cLang.PluginsArticleNameHasUsedCode, e.Lang)
	}

	//确保分类存在
	categoryService := NewContentCategoryService(&e.Service)
	_, respCode, err = categoryService.Get(c.CateId, nil)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if err != nil && respCode == lang.DataNotFoundCode {
		return 0, cLang.PluginsCategoryNotFoundCode, lang.MsgErr(cLang.PluginsCategoryNotFoundCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update
// @Description: 修改ContentArticle对象
// @receiver e
// @param c
// @param p
// @return bool 是否有数据更新
// @return error
func (e *ContentArticle) Update(c *dto.ContentArticleUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return false, cLang.PluginsArticleNameCode, lang.MsgErr(cLang.PluginsArticleNameCode, e.Lang)
	}
	if c.Content == "" {
		return false, cLang.PluginsArticleContentCode, lang.MsgErr(cLang.PluginsArticleContentCode, e.Lang)
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
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDictDataValueExistCode, lang.MsgErr(sysLang.SysDictDataValueExistCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// Remove
// @Description: 删除ContentArticle
// @receiver e
// @param ids
// @param p
// @return int
// @return error
func (e *ContentArticle) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.ContentArticle
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出ContentArticle excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *ContentArticle) GetExcel(list []models.ContentArticle) ([]byte, error) {
	sheetName := "ContentArticle"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	xlsx.SetColWidth(sheetName, "A", "P", 25)
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"文章编号", "分类名称", "标题", "内容", "备注", "时间"})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)

		//按标签对应输入数据
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.ContentCategory.Name, item.Name, item.Content, item.Remark, dateUtils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
