package service

import (
	"go-admin/app/admin/sys/constant"
	sysLang "go-admin/app/admin/sys/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm/clause"
	"time"

	"github.com/casbin/casbin/v2"

	"gorm.io/gorm"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysRole struct {
	service.Service
}

// NewSysRoleService admin-实例化角色管理
func NewSysRoleService(s *service.Service) *SysRole {
	var srv = new(SysRole)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetList admin-获取角色管理全部列表
func (e *SysRole) GetList(c *dto.SysRoleQueryReq, p *middleware.DataPermission) ([]models.SysRole, int64, int, error) {
	var list []models.SysRole
	var data models.SysRole
	var count int64

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// GetPage admin-获取角色管理分页列表
func (e *SysRole) GetPage(c *dto.SysRoleQueryReq, p *middleware.DataPermission) ([]models.SysRole, int64, int, error) {
	var list []models.SysRole
	var data models.SysRole
	var count int64

	err := e.Orm.Model(&data).Preload("SysMenu").
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

// Get admin-获取角色管理详情
func (e *SysRole) Get(id int64, p *middleware.DataPermission) (*models.SysRole, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysRole{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}

	menuIds, respCode, err := e.GetMenuIdsByRole(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	deptIds, respCode, err := e.GetDeptIdsByRole(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	data.MenuIds = menuIds
	data.DeptIds = deptIds
	return data, lang.SuccessCode, nil
}

// QueryOne admin-获取角色管理一条记录
func (e *SysRole) QueryOne(queryCondition *dto.SysRoleQueryReq, p *middleware.DataPermission) (*models.SysRole, int, error) {
	data := &models.SysRole{}
	err := e.Orm.Model(&models.SysRole{}).
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

// Count admin-获取角色管理数据总数
func (e *SysRole) Count(c *dto.SysRoleQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysRole{}).
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

// Insert admin-创建角色管理
func (e *SysRole) Insert(c *dto.SysRoleInsertReq, cb *casbin.SyncedEnforcer) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	if c.RoleName == "" {
		return 0, sysLang.SysRoleNameEmptyCode, lang.MsgErr(sysLang.SysRoleNameEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, sysLang.SysRoleStatusEmptyCode, lang.MsgErr(sysLang.SysRoleStatusEmptyCode, e.Lang)
	}
	if c.RoleKey == "" {
		return 0, sysLang.SysRoleKeyEmptyCode, lang.MsgErr(sysLang.SysRoleKeyEmptyCode, e.Lang)
	}
	if c.RoleSort < 0 {
		return 0, sysLang.SysRoleSortEmptyCode, lang.MsgErr(sysLang.SysRoleSortEmptyCode, e.Lang)
	}

	//确保角色key不存在
	req := dto.SysRoleQueryReq{}
	req.RoleKey = c.RoleKey
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysRoleKeyExistCode, lang.MsgErr(sysLang.SysRoleKeyExistCode, e.Lang)
	}

	//获取菜单
	sysMenuService := NewSysMenuService(&e.Service)
	sysMenuQueryReq := dto.SysMenuQueryReq{}
	sysMenuQueryReq.MenuIds = c.MenuIds
	sysMens, respCode, err := sysMenuService.GetList(&sysMenuQueryReq, true)
	if err != nil {
		return 0, respCode, err
	}

	//插入数据 菜单和权限
	now := time.Now()
	data := models.SysRole{}
	data.RoleName = c.RoleName
	data.Status = c.Status
	data.RoleKey = c.RoleKey
	data.RoleSort = c.RoleSort
	data.Remark = c.Remark
	data.DataScope = c.DataScope
	data.SysMenu = &sysMens
	data.SysDept = c.SysDept
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now

	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	err = e.Orm.Save(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}

	//casbin
	for _, menu := range sysMens {
		for _, api := range menu.SysApi {
			_, err = cb.AddNamedPolicy("p", data.RoleKey, api.Path, api.Method)
		}
	}
	_ = cb.SavePolicy()
	return data.Id, lang.SuccessCode, nil
}

// Update admin-更新角色管理
func (e *SysRole) Update(c *dto.SysRoleUpdateReq, cb *casbin.SyncedEnforcer) (int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error

	//确保角色key不存在
	req := dto.SysRoleQueryReq{}
	req.RoleKey = c.RoleKey
	role, respCode, err := e.QueryOne(&req, nil)
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if respCode == lang.SuccessCode && role.Id != c.Id {
		return sysLang.SysRoleKeyExistCode, lang.MsgErr(sysLang.SysRoleKeyExistCode, e.Lang)
	}
	if role != nil && role.RoleKey == constant.RoleKeyAdmin {
		return sysLang.SysRoleAdminNoOpCode, lang.MsgErr(sysLang.SysRoleAdminNoOpCode, e.Lang)
	}

	e.Orm = e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	//删除角色对应的菜单和api
	var data = models.SysRole{}
	var mlist = make([]models.SysMenu, 0)
	e.Orm.Preload("SysMenu").First(&data, c.Id)
	e.Orm.Preload("SysApi").Where("id in ?", c.MenuIds).Find(&mlist)
	err = e.Orm.Model(&data).Association("SysMenu").Delete(data.SysMenu)
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}

	//更新角色信息
	now := time.Now()
	data.RoleName = c.RoleName
	data.Status = c.Status
	data.RoleKey = c.RoleKey
	data.RoleSort = c.RoleSort
	data.Remark = c.Remark
	data.DataScope = c.DataScope
	data.SysDept = c.SysDept
	data.SysMenu = &mlist
	data.UpdatedAt = &now
	data.UpdateBy = c.CurrUserId
	err = e.Orm.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&data).Error
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}

	//casbin
	_, err = cb.RemoveFilteredPolicy(0, data.RoleKey)
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	for _, menu := range mlist {
		for _, api := range menu.SysApi {
			_, err = cb.AddNamedPolicy("p", data.RoleKey, api.Path, api.Method)
		}
	}
	_ = cb.SavePolicy()
	return lang.SuccessCode, nil
}

// Delete admin-删除角色管理
func (e *SysRole) Delete(ids []int64) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userService := NewSysUserService(&e.Service)
	userReq := dto.SysUserQueryReq{}
	userReq.RoleIds = ids
	count, respCode, err := userService.Count(&userReq)
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return sysLang.SysRoleUserExistNoDeleteCode, lang.MsgErr(sysLang.SysRoleUserExistNoDeleteCode, e.Lang)
	}

	for _, id := range ids {
		var role = models.SysRole{}
		err = tx.Preload("SysMenu").Preload("SysDept").First(&role, id).Error
		if err != nil {
			return lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
		}

		if role.RoleKey == constant.RoleKeyAdmin {
			return sysLang.SysRoleAdminNoOpCode, lang.MsgErr(sysLang.SysRoleAdminNoOpCode, e.Lang)
		}
		err = tx.Select(clause.Associations).Delete(&role).Error
		if err != nil {
			return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
		}
	}
	return lang.SuccessCode, nil
}

// GetMenuIdsByRole admin-获取角色对应的菜单编号集合
func (e *SysRole) GetMenuIdsByRole(roleId int64) ([]int64, int, error) {
	if roleId <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	menuIds := make([]int64, 0)
	model := models.SysRole{}
	model.Id = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].Id)
	}
	return menuIds, lang.SuccessCode, nil
}

// GetDeptIdsByRole admin-获取角色对应的的部门编号集合
func (e *SysRole) GetDeptIdsByRole(roleId int64) ([]int64, int, error) {
	deptIds := make([]int64, 0)
	deptList := make([]dto.SysRoleDeptResp, 0)
	if err := e.Orm.Table("admin_sys_role_dept").
		Select("admin_sys_role_dept.dept_id").
		Joins("LEFT JOIN admin_sys_dept on admin_sys_dept.id=admin_sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" admin_sys_role_dept.dept_id not in(select admin_sys_dept.parent_id from admin_sys_role_dept LEFT JOIN admin_sys_dept on admin_sys_dept.id=admin_sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, lang.SuccessCode, nil
}

// UpdateDataScope admin-更新角色管理数据权限
func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) (int, error) {
	var err error
	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()
	var dlist = make([]models.SysDept, 0)
	var model = models.SysRole{}
	e.Orm.Preload("SysDept").First(&model, c.Id)                           //查找角色id所属部门
	e.Orm.Where("id in ?", c.DeptIds).Find(&dlist)                         //查找所选部门id对应的部门信息
	err = e.Orm.Model(&model).Association("SysDept").Delete(model.SysDept) //删除角色原有的部门信息
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	if c.Id > 0 {
		model.Id = c.Id
	}
	model.DataScope = c.DataScope
	model.DeptIds = c.DeptIds
	model.SysDept = dlist
	err = e.Orm.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model).Error
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	return lang.SuccessCode, nil
}

// UpdateStatus admin-更新角色管理状态
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) (int, error) {
	var err error
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	err = tx.First(&model, c.RoleId).Error
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	model.Status = c.Status
	err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model).Error
	if err != nil {
		return lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetWithName admin-根据角色名获取角色详情
func (e *SysRole) GetWithName(d *dto.SysRoleQueryReq) (*models.SysRole, int, error) {
	model := &models.SysRole{}
	err := e.Orm.Where("role_name = ?", d.RoleName).First(model).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	menuIds, respCode, err := e.GetMenuIdsByRole(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	deptIds, respCode, err := e.GetDeptIdsByRole(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	model.MenuIds = menuIds
	model.DeptIds = deptIds
	return model, lang.SuccessCode, nil
}

// GetPermissionsByRoleId admin-根据角色获取权限
func (e *SysRole) GetPermissionsByRoleId(roleId int64) ([]string, int, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.Id = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, lang.SuccessCode, nil
}
