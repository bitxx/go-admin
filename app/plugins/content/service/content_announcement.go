package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	cLang "go-admin/app/plugins/content/lang"
	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm"
	"time"
)

type ContentAnnouncement struct {
	service.Service
}

// NewContentAnnouncementService plugins-实例化公告管理
func NewContentAnnouncementService(s *service.Service) *ContentAnnouncement {
	var srv = new(ContentAnnouncement)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage plugins-获取公告管理分页列表
func (e *ContentAnnouncement) GetPage(c *dto.ContentAnnouncementQueryReq, p *middleware.DataPermission) ([]models.ContentAnnouncement, int64, int, error) {
	var data models.ContentAnnouncement
	var list []models.ContentAnnouncement
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

// Get plugins-获取公告管理详情
func (e *ContentAnnouncement) Get(id int64, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.ContentAnnouncement{}
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

// QueryOne plugins-获取公告管理一条记录
func (e *ContentAnnouncement) QueryOne(queryCondition *dto.ContentAnnouncementQueryReq, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	data := &models.ContentAnnouncement{}
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

// Count admin-获取公告管理数据总数
func (e *ContentAnnouncement) Count(queryCondition *dto.ContentAnnouncementQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentAnnouncement{}).
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

// Insert plugins-新增公告管理
func (e *ContentAnnouncement) Insert(c *dto.ContentAnnouncementInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Title == "" {
		return 0, cLang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(cLang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return 0, cLang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(cLang.PluginsAnnouncementContentEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, cLang.PluginsAnnouncementStatusEmptyCode, lang.MsgErr(cLang.PluginsAnnouncementStatusEmptyCode, e.Lang)
	}
	if c.Num < 0 {
		return 0, cLang.PluginsAnnouncementNumCode, lang.MsgErr(cLang.PluginsAnnouncementNumCode, e.Lang)
	}
	now := time.Now()
	var data models.ContentAnnouncement
	data.Title = c.Title
	data.Content = c.Content
	data.Num = c.Num
	data.Remark = c.Remark
	data.Status = c.Status
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	data.CreatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update plugins-更新公告管理
func (e *ContentAnnouncement) Update(c *dto.ContentAnnouncementUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Num < 0 {
		return false, cLang.PluginsAnnouncementNumCode, lang.MsgErr(cLang.PluginsAnnouncementNumCode, e.Lang)
	}
	if c.Title == "" {
		return false, cLang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(cLang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return false, cLang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(cLang.PluginsAnnouncementContentEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.Title != "" && data.Title != c.Title {
		req := dto.ContentAnnouncementQueryReq{}
		req.Title = c.Title
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, cLang.PluginsAnnouncementTitleHasUsedCode, lang.MsgErr(cLang.PluginsAnnouncementTitleHasUsedCode, e.Lang)
		}
		updates["title"] = c.Title
	}
	if c.Content != "" && data.Content != c.Content {
		updates["content"] = c.Content
	}
	if c.Num > 0 && data.Num != c.Num {
		updates["num"] = c.Num
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

// Delete plugins-删除公告管理
func (e *ContentAnnouncement) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.ContentAnnouncement
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// Export plugins-导出公告管理
func (e *ContentAnnouncement) Export(list []models.ContentAnnouncement) ([]byte, error) {
	sheetName := "ContentAnnouncement"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "E", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"公告编号", "标题", "内容", "阅读次数", "备注信息"})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.Title, item.Content, item.Num, item.Remark,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
