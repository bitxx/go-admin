package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	sysLang "go-admin/app/admin/lang"
	"go-admin/common/dto/service"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type SysPost struct {
	service.Service
}

func NewSysPostService(s *service.Service) *SysPost {
	var srv = new(SysPost)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysPost列表
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get 获取SysPost对象
func (e *SysPost) Get(id int64, p *middleware.DataPermission) (*models.SysPost, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysPost{}
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

// QueryOne 通过自定义条件获取一条记录
func (e *SysPost) QueryOne(queryCondition *dto.SysPostQueryReq, p *middleware.DataPermission) (*models.SysPost, int, error) {
	data := &models.SysPost{}
	err := e.Orm.Model(&models.SysPost{}).
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

// Count 获取条数
func (e *SysPost) Count(c *dto.SysPostQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysPost{}).
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

// Insert 创建SysPost对象
func (e *SysPost) Insert(c *dto.SysPostInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.PostName == "" {
		return 0, sysLang.SysPostNameEmptyCode, lang.MsgErr(sysLang.SysPostNameEmptyCode, e.Lang)
	}
	if c.PostCode == "" {
		return 0, sysLang.SysPostCodeEmptyCode, lang.MsgErr(sysLang.SysPostCodeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return 0, sysLang.SysPostSortEmptyCode, lang.MsgErr(sysLang.SysPostSortEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, sysLang.SysPostStatusEmptyCode, lang.MsgErr(sysLang.SysPostStatusEmptyCode, e.Lang)
	}

	//确保岗位名称不存在
	req := dto.SysPostQueryReq{}
	req.PostName = c.PostName
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysPostNameExistCode, lang.MsgErr(sysLang.SysPostNameExistCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改SysPost对象
func (e *SysPost) Update(c *dto.SysPostUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.PostName == "" {
		return false, sysLang.SysPostNameEmptyCode, lang.MsgErr(sysLang.SysPostNameEmptyCode, e.Lang)
	}
	if c.PostCode == "" {
		return false, sysLang.SysPostCodeEmptyCode, lang.MsgErr(sysLang.SysPostCodeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return false, sysLang.SysPostSortEmptyCode, lang.MsgErr(sysLang.SysPostSortEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return false, sysLang.SysPostStatusEmptyCode, lang.MsgErr(sysLang.SysPostStatusEmptyCode, e.Lang)
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
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysPostNameExistCode, lang.MsgErr(sysLang.SysPostNameExistCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// Remove 删除SysPost
func (e *SysPost) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	for _, id := range ids {
		userService := NewSysUserService(&e.Service)
		userReq := dto.SysUserQueryReq{}
		userReq.RoleId = id
		count, respCode, err := userService.Count(&userReq)
		if err != nil && respCode != lang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return sysLang.SysRoleUserExistNoDeleteCode, lang.MsgErr(sysLang.SysRoleUserExistNoDeleteCode, e.Lang)
		}
	}

	var err error
	var data models.SysPost
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel 导出SysPost
func (e *SysPost) GetExcel(list []models.SysPost) ([]byte, error) {
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
		postStatus := dictService.GetLabel("sys_status", item.Status)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.PostName, item.PostCode, item.Sort, postStatus, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
