package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"

	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysPost struct {
	service.Service
}

// NewSysPostService admin-实例化岗位管理
func NewSysPostService(s *service.Service) *SysPost {
	var srv = new(SysPost)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTotalList admin-获取岗位管理全部列表
func (e *SysPost) GetTotalList(c *dto.SysPostQueryReq, p *middleware.DataPermission) ([]models.SysPost, int64, int, error) {
	var list []models.SysPost
	var data models.SysPost
	var count int64

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// GetPage admin-获取岗位管理分页列表
func (e *SysPost) GetPage(c *dto.SysPostQueryReq, p *middleware.DataPermission) ([]models.SysPost, int64, int, error) {
	var list []models.SysPost
	var data models.SysPost
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

// Get admin-获取岗位管理详情
func (e *SysPost) Get(id int64, p *middleware.DataPermission) (*models.SysPost, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysPost{}
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

// QueryOne admin-获取岗位管理一条记录
func (e *SysPost) QueryOne(queryCondition *dto.SysPostQueryReq, p *middleware.DataPermission) (*models.SysPost, int, error) {
	data := &models.SysPost{}
	err := e.Orm.Model(&models.SysPost{}).
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

// Count admin-获取岗位管理数据总数
func (e *SysPost) Count(c *dto.SysPostQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysPost{}).
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

// Insert admin-新增岗位管理
func (e *SysPost) Insert(c *dto.SysPostInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.PostName == "" {
		return 0, baseLang.SysPostNameEmptyCode, lang.MsgErr(baseLang.SysPostNameEmptyCode, e.Lang)
	}
	if c.PostCode == "" {
		return 0, baseLang.SysPostCodeEmptyCode, lang.MsgErr(baseLang.SysPostCodeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return 0, baseLang.SysPostSortEmptyCode, lang.MsgErr(baseLang.SysPostSortEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, baseLang.SysPostStatusEmptyCode, lang.MsgErr(baseLang.SysPostStatusEmptyCode, e.Lang)
	}

	//确保岗位名称不存在
	req := dto.SysPostQueryReq{}
	req.PostName = c.PostName
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.SysPostNameExistCode, lang.MsgErr(baseLang.SysPostNameExistCode, e.Lang)
	}

	now := time.Now()
	data := models.SysPost{}
	data.PostName = c.PostName
	data.PostCode = c.PostCode
	data.Sort = c.Sort
	data.Status = c.Status
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

// Update admin-更新岗位管理
func (e *SysPost) Update(c *dto.SysPostUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.PostName == "" {
		return false, baseLang.SysPostNameEmptyCode, lang.MsgErr(baseLang.SysPostNameEmptyCode, e.Lang)
	}
	if c.PostCode == "" {
		return false, baseLang.SysPostCodeEmptyCode, lang.MsgErr(baseLang.SysPostCodeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return false, baseLang.SysPostSortEmptyCode, lang.MsgErr(baseLang.SysPostSortEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return false, baseLang.SysPostStatusEmptyCode, lang.MsgErr(baseLang.SysPostStatusEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}

	if c.PostName != "" && data.PostName != c.PostName {
		updates["post_name"] = c.PostName
	}
	if c.PostCode != "" && data.PostCode != c.PostCode {
		//判断岗位名称是否已存在
		req := dto.SysPostQueryReq{}
		req.PostCode = c.PostCode
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysPostNameExistCode, lang.MsgErr(baseLang.SysPostNameExistCode, e.Lang)
		}
		updates["post_code"] = c.PostCode
	}
	if c.Sort > 0 && data.Sort != c.Sort {
		updates["sort"] = c.Sort
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
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

// Delete admin-删除岗位管理
func (e *SysPost) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	for _, id := range ids {
		userService := NewSysUserService(&e.Service)
		userReq := dto.SysUserQueryReq{}
		userReq.RoleId = id
		count, respCode, err := userService.Count(&userReq)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return baseLang.SysRoleUserExistNoDeleteCode, lang.MsgErr(baseLang.SysRoleUserExistNoDeleteCode, e.Lang)
		}
	}

	var err error
	var data models.SysPost
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export admin-导出岗位管理
func (e *SysPost) Export(list []models.SysPost) ([]byte, error) {
	//sheet名称
	sheetName := "Post"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "F", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"岗位编号", "岗位名称", "岗位编码", "岗位排序", "状态", "创建时间"})
	dictService := NewSysDictDataService(&e.Service)

	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		postStatus := dictService.GetLabel("admin_sys_status", item.Status)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.PostName, item.PostCode, item.Sort, postStatus, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
