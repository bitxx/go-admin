package service

import (
	"errors"
	"go-admin/internal/app/admin/sys/models"
	"go-admin/internal/app/admin/sys/service/dto"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/tree"
	"gorm.io/gorm"
	"strconv"
	"time"

	cDto "go-admin/pkg/dto"
)

type SysDept struct {
	service.Service
}

// NewSysDeptService admin-实例化部门管理
func NewSysDeptService(s *service.Service) *SysDept {
	var srv = new(SysDept)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTreeList admin-获取部门树列表
func (e *SysDept) GetTreeList(c *dto.SysDeptQueryReq) ([]*models.SysDept, int, error) {
	list, respCode, err := e.getList(c)
	if err != nil {
		return nil, respCode, err
	}
	treeList := tree.GenTree(&list,
		func(item models.SysDept) int64 { return item.Id },
		func(item models.SysDept) int64 { return item.ParentId },
		func(item *models.SysDept, children []*models.SysDept) { item.Children = children },
	)

	return []*models.SysDept{
		{Id: 0, DeptName: "主目录", ParentId: 0, Children: treeList},
	}, clang.SuccessCode, nil
}

// Get admin-获取部门管理详情
func (e *SysDept) Get(id int64, p *middleware.DataPermission) (*models.SysDept, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.SysDept{}
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

// QueryOne admin-获取部门管理一条记录
func (e *SysDept) QueryOne(queryCondition *dto.SysDeptQueryReq, p *middleware.DataPermission) (*models.SysDept, int, error) {
	data := &models.SysDept{}
	err := e.Orm.Model(&models.SysDept{}).
		Scopes(
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

// Count admin-获取部门管理数据总数
func (e *SysDept) Count(c *dto.SysDeptQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDept{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return count, clang.SuccessCode, nil
}

// Insert admin-新增部门管理
func (e *SysDept) Insert(c *dto.SysDeptInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.ParentId.IntPart() < 0 {
		return 0, clang.SysDeptParentIdEmptyCode, lang.MsgErr(clang.SysDeptParentIdEmptyCode, e.Lang)
	}
	if c.DeptName == "" {
		return 0, clang.SysDeptNameEmptyCode, lang.MsgErr(clang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return 0, clang.SysDeptLeaderEmptyCode, lang.MsgErr(clang.SysDeptLeaderEmptyCode, e.Lang)
	}

	//确保部门名称不存在
	req := dto.SysDeptQueryReq{}
	req.DeptName = c.DeptName
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != clang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, clang.SysDeptNameExistCode, lang.MsgErr(clang.SysDeptNameExistCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update admin-更新部门管理
func (e *SysDept) Update(c *dto.SysDeptUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.DeptName == "" {
		return false, clang.SysDeptNameEmptyCode, lang.MsgErr(clang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return false, clang.SysDeptLeaderEmptyCode, lang.MsgErr(clang.SysDeptLeaderEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	if c.Id == c.ParentId {
		return false, clang.SysDeptParentSelfCode, lang.MsgErr(clang.SysDeptParentSelfCode, e.Lang)
	}

	updates := map[string]interface{}{}
	if c.DeptName != "" && data.DeptName != c.DeptName {
		req := dto.SysDeptQueryReq{}
		req.DeptName = c.DeptName
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysDeptNameExistCode, lang.MsgErr(clang.SysDeptNameExistCode, e.Lang)
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
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete admin-删除部门管理
func (e *SysDept) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	for _, id := range ids {
		//若被使用，不得删除
		dataReq := dto.SysDeptQueryReq{}
		dataReq.ParentIds = "," + strconv.FormatInt(id, 10) + ","
		count, respCode, err := e.Count(&dataReq)
		if err != nil && respCode != clang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return clang.SysDeptChildExistNoDelCode, lang.MsgErr(clang.SysDeptChildExistNoDelCode, e.Lang)
		}
	}

	var err error
	var data models.SysDept
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
}

// getList admin-获取部门管理全部列表
func (e *SysDept) getList(c *dto.SysDeptQueryReq) ([]models.SysDept, int, error) {
	var list []models.SysDept
	err := e.Orm.Model(&models.SysDept{}).Order("sort").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, clang.SuccessCode, nil
}
