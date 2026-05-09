package service

import (
	"errors"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/xuri/excelize/v2"
	"go-admin/internal/app/plugins/content/models"
	"go-admin/internal/app/plugins/content/service/dto"
	clang "go-admin/internal/common/lang"
	cDto "go-admin/pkg/dto"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
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
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// Get plugins-获取公告管理详情
func (e *ContentAnnouncement) Get(id int64, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.ContentAnnouncement{}
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

// QueryOne plugins-获取公告管理一条记录
func (e *ContentAnnouncement) QueryOne(queryCondition *dto.ContentAnnouncementQueryReq, p *middleware.DataPermission) (*models.ContentAnnouncement, int, error) {
	data := &models.ContentAnnouncement{}
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

// Count admin-获取公告管理数据总数
func (e *ContentAnnouncement) Count(queryCondition *dto.ContentAnnouncementQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.ContentAnnouncement{}).
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

// Insert plugins-新增公告管理
func (e *ContentAnnouncement) Insert(c *dto.ContentAnnouncementInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Title == "" {
		return 0, clang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(clang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return 0, clang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(clang.PluginsAnnouncementContentEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, clang.PluginsAnnouncementStatusEmptyCode, lang.MsgErr(clang.PluginsAnnouncementStatusEmptyCode, e.Lang)
	}
	if c.Num < 0 {
		return 0, clang.PluginsAnnouncementNumCode, lang.MsgErr(clang.PluginsAnnouncementNumCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update plugins-更新公告管理
func (e *ContentAnnouncement) Update(c *dto.ContentAnnouncementUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Num < 0 {
		return false, clang.PluginsAnnouncementNumCode, lang.MsgErr(clang.PluginsAnnouncementNumCode, e.Lang)
	}
	if c.Title == "" {
		return false, clang.PluginsAnnouncementTitleEmptyCode, lang.MsgErr(clang.PluginsAnnouncementTitleEmptyCode, e.Lang)
	}
	if c.Content == "" {
		return false, clang.PluginsAnnouncementContentEmptyCode, lang.MsgErr(clang.PluginsAnnouncementContentEmptyCode, e.Lang)
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
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.PluginsAnnouncementTitleHasUsedCode, lang.MsgErr(clang.PluginsAnnouncementTitleHasUsedCode, e.Lang)
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
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete plugins-删除公告管理
func (e *ContentAnnouncement) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	var data models.ContentAnnouncement
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
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
