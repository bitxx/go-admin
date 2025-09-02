package service

import (
	"errors"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/xuri/excelize/v2"
	"go-admin/app/plugins/content/models"
	"go-admin/app/plugins/content/service/dto"
	baseLang "go-admin/config/base/lang"

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
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get plugins-获取公告管理详情
func (e *ContentAnnouncement) Get(id int64, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.ContentAnnouncement{}
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

// QueryOne plugins-获取公告管理一条记录
func (e *ContentAnnouncement) QueryOne(queryCondition *dto.ContentAnnouncementQueryReq, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	data := &models.ContentAnnouncement{}
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

// Count admin-获取公告管理数据总数
func (e *ContentAnnouncement) Count(queryCondition *dto.ContentAnnouncementQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentAnnouncement{}).
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

// Insert plugins-新增公告管理
func (e *ContentAnnouncement) Insert(c *dto.ContentAnnouncementInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Title == "" {
		return 0, baseLang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(baseLang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return 0, baseLang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(baseLang.PluginsAnnouncementContentEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, baseLang.PluginsAnnouncementStatusEmptyCode, lang.MsgErr(baseLang.PluginsAnnouncementStatusEmptyCode, e.Lang)
	}
	if c.Num < 0 {
		return 0, baseLang.PluginsAnnouncementNumCode, lang.MsgErr(baseLang.PluginsAnnouncementNumCode, e.Lang)
	}
	now := time.Now()
	var data models.ContentAnnouncement
	data.Title = c.Title
	data.Content = bluemonday.UGCPolicy().Sanitize(c.Content)
	data.Num = c.Num
	data.Remark = c.Remark
	data.Status = c.Status
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	data.CreatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update plugins-更新公告管理
func (e *ContentAnnouncement) Update(c *dto.ContentAnnouncementUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Num < 0 {
		return false, baseLang.PluginsAnnouncementNumCode, lang.MsgErr(baseLang.PluginsAnnouncementNumCode, e.Lang)
	}
	if c.Title == "" {
		return false, baseLang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(baseLang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return false, baseLang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(baseLang.PluginsAnnouncementContentEmptyCode, e.Lang)
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
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.PluginsAnnouncementTitleHasUsedCode, lang.MsgErr(baseLang.PluginsAnnouncementTitleHasUsedCode, e.Lang)
		}
		updates["title"] = c.Title
	}
	if c.Content != "" && data.Content != c.Content {
		updates["content"] = bluemonday.UGCPolicy().Sanitize(c.Content)
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
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete plugins-删除公告管理
func (e *ContentAnnouncement) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.ContentAnnouncement
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
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
