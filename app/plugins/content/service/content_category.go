package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	sysLang "go-admin/app/admin/lang"
	cLang "go-admin/app/plugins/content/lang"
	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
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

// NewContentCategoryService
// @Description: 实例化ContentCategory
// @param s
// @return *ContentCategory
func NewContentCategoryService(s *service.Service) *ContentCategory {
	var srv = new(ContentCategory)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取ContentCategory列表
// @receiver e
// @param c
// @param p
// @return []models.ContentCategory
// @return int64
// @return int
// @return error
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get
// @Description: 获取ContentCategory对象
// @receiver e
// @param id 编号
// @param p
// @return *models.ContentCategory
// @return int
// @return error
func (e *ContentCategory) Get(id int64, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.ContentCategory{}
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
// @Description: 通过自定义条件获取ContentCategory一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.ContentCategory
// @return error
func (e *ContentCategory) QueryOne(queryCondition *dto.ContentCategoryQueryReq, p *middleware.DataPermission) (*models.ContentCategory, int, error) {
	data := &models.ContentCategory{}
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
func (e *ContentCategory) Count(queryCondition *dto.ContentCategoryQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentCategory{}).
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
// @Description: 创建ContentCategory对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *ContentCategory) Insert(c *dto.ContentCategoryInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return 0, cLang.PluginsCategoryNameCode, lang.MsgErr(cLang.PluginsCategoryNameCode, e.Lang)
	}
	req := dto.ContentCategoryQueryReq{}
	req.NameInner = c.Name
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, cLang.PluginsCategoryNameHasUsedCode, lang.MsgErr(cLang.PluginsCategoryNameHasUsedCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update
// @Description: 修改ContentCategory对象
// @receiver e
// @param c
// @param p
// @return bool 是否有数据更新
// @return error
func (e *ContentCategory) Update(c *dto.ContentCategoryUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Name == "" {
		return false, cLang.PluginsCategoryNameCode, lang.MsgErr(cLang.PluginsCategoryNameCode, e.Lang)
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
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDictDataValueExistCode, lang.MsgErr(sysLang.SysDictDataValueExistCode, e.Lang)
		}
		updates["name"] = c.Name
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
// @Description: 删除ContentCategory
// @receiver e
// @param ids
// @param p
// @return int
// @return error
func (e *ContentCategory) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	//若有文章，不得删除
	articleService := NewContentArticleService(&e.Service)
	articleReq := dto.ContentArticleQueryReq{}
	articleReq.CateIds = ids
	count, respCode, err := articleService.Count(&articleReq)
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return cLang.PluginsCategoryNameHasUsedCode, lang.MsgErr(cLang.PluginsCategoryNameHasUsedCode, e.Lang)
	}
	var data models.ContentCategory
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出ContentCategory excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *ContentCategory) GetExcel(list []models.ContentCategory) ([]byte, error) {
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
