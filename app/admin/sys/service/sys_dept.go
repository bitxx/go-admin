package service

import (
	sysLang "go-admin/app/admin/sys/lang"
	"go-admin/app/admin/sys/models"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/tree"
	"gorm.io/gorm"
	"strconv"
	"time"

	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysDept struct {
	service.Service
}

// NewSysDeptService sys-实例化部门管理
func NewSysDeptService(s *service.Service) *SysDept {
	var srv = new(SysDept)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTreeList sys-获取部门树列表
func (e *SysDept) GetTreeList(c *dto.SysDeptQueryReq) ([]*models.SysDept, int, error) {
	list, respCode, err := e.getList(c)
	if err != nil {
		return nil, respCode, err
	}
	return tree.GenTree(&list,
		func(item models.SysDept) int64 { return item.Id },
		func(item models.SysDept) int64 { return item.ParentId },
		func(item *models.SysDept, children []*models.SysDept) { item.Children = children },
	), lang.SuccessCode, nil
}

// Get sys-获取部门管理详情
func (e *SysDept) Get(id int64, p *middleware.DataPermission) (*models.SysDept, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysDept{}
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

// QueryOne sys-获取部门管理一条记录
func (e *SysDept) QueryOne(queryCondition *dto.SysDeptQueryReq, p *middleware.DataPermission) (*models.SysDept, int, error) {
	data := &models.SysDept{}
	err := e.Orm.Model(&models.SysDept{}).
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

// Count sys-获取部门管理数据总数
func (e *SysDept) Count(c *dto.SysDeptQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDept{}).
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

// Insert sys-添加部门管理
func (e *SysDept) Insert(c *dto.SysDeptInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.ParentId.IntPart() < 0 {
		return 0, sysLang.SysDeptParentIdEmptyCode, lang.MsgErr(sysLang.SysDeptParentIdEmptyCode, e.Lang)
	}
	if c.DeptName == "" {
		return 0, sysLang.SysDeptNameEmptyCode, lang.MsgErr(sysLang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return 0, sysLang.SysDeptLeaderEmptyCode, lang.MsgErr(sysLang.SysDeptLeaderEmptyCode, e.Lang)
	}

	//确保部门名称不存在
	req := dto.SysDeptQueryReq{}
	req.DeptName = c.DeptName
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysDeptNameExistCode, lang.MsgErr(sysLang.SysDeptNameExistCode, e.Lang)
	}
	//path组装
	parentIds := "0,"
	if c.ParentId.IntPart() > 0 {
		dept := &models.SysDept{}
		dept, respCode, err = e.Get(c.ParentId.IntPart(), nil)
		if err != nil {
			return 0, respCode, err
		}
		parentIds = dept.ParentIds + "," + strconv.FormatInt(dept.Id, 10) + ","
	}

	now := time.Now()
	data := models.SysDept{}
	data.DeptName = c.DeptName
	data.ParentId = c.ParentId.IntPart()
	data.ParentIds = parentIds
	data.Sort = c.Sort
	data.Leader = c.Leader
	data.Phone = c.Phone
	data.Email = c.Email
	data.Status = global.SysStatusOk
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

// Update sys-更新部门管理
func (e *SysDept) Update(c *dto.SysDeptUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.DeptName == "" {
		return false, sysLang.SysDeptNameEmptyCode, lang.MsgErr(sysLang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return false, sysLang.SysDeptLeaderEmptyCode, lang.MsgErr(sysLang.SysDeptLeaderEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	if c.Id == c.ParentId {
		return false, sysLang.SysDeptParentSelfCode, lang.MsgErr(sysLang.SysDeptParentSelfCode, e.Lang)
	}

	updates := map[string]interface{}{}
	if c.DeptName != "" && data.DeptName != c.DeptName {
		req := dto.SysDeptQueryReq{}
		req.DeptName = c.DeptName
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDeptNameExistCode, lang.MsgErr(sysLang.SysDeptNameExistCode, e.Lang)
		}
		updates["dept_name"] = c.DeptName
	}
	//
	if c.Sort > 0 && data.Sort != c.Sort {
		updates["sort"] = c.Sort
	}
	if c.ParentId > 0 && data.ParentId != c.ParentId {
		updates["parent_id"] = c.ParentId
	}
	if c.Leader != "" && data.Leader != c.Leader {
		updates["leader"] = c.Leader
	}
	if c.Phone != "" && data.Phone != c.Phone {
		updates["phone"] = c.Phone
	}
	if c.Email != "" && data.Email != c.Email {
		updates["email"] = c.Email
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

// Delete sys-删除部门管理
func (e *SysDept) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	for _, id := range ids {
		//若被使用，不得删除
		dataReq := dto.SysDeptQueryReq{}
		dataReq.ParentIds = "," + strconv.FormatInt(id, 10) + ","
		count, respCode, err := e.Count(&dataReq)
		if err != nil && respCode != lang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return sysLang.SysDeptChildExistNoDelCode, lang.MsgErr(sysLang.SysDeptChildExistNoDelCode, e.Lang)
		}
	}

	var err error
	var data models.SysDept
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// getList sys-获取部门管理全部列表
func (e *SysDept) getList(c *dto.SysDeptQueryReq) ([]models.SysDept, int, error) {
	var list []models.SysDept
	err := e.Orm.Model(&models.SysDept{}).Order("sort").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}
