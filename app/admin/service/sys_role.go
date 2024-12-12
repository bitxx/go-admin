package service

import (
	"go-admin/app/admin/constant"
	sysLang "go-admin/app/admin/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm/clause"
	"time"

	"github.com/casbin/casbin/v2"

	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/core/dto"
)

type SysRole struct {
	service.Service
}

func NewSysRoleService(s *service.Service) *SysRole {
	var srv = new(SysRole)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTotalList 获取SysPost列表
func (e *SysRole) GetTotalList(c *dto.SysRoleQueryReq, p *middleware.DataPermission) ([]models.SysRole, int64, int, error) {
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

// GetPage 获取SysRole列表
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

// Get 获取SysRole对象
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

	menuIds, respCode, err := e.GetRoleMenuId(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	deptIds, respCode, err := e.GetRoleDeptId(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	data.MenuIds = menuIds
	data.DeptIds = deptIds
	return data, lang.SuccessCode, nil
}

// QueryOne 通过自定义条件获取一条记录
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

// Count 获取条数
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

// Insert 创建SysRole对象
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

// Update 修改SysRole对象
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

// Remove 删除SysRole
func (e *SysRole) Remove(ids []int64) (int, error) {
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

// GetRoleMenuId 获取角色对应的菜单ids
func (e *SysRole) GetRoleMenuId(roleId int64) ([]int64, int, error) {
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

// GetRoleDeptId 获取角色的部门ID集合
func (e *SysRole) GetRoleDeptId(roleId int64) ([]int64, int, error) {
	deptIds := make([]int64, 0)
	deptList := make([]dto.SysRoleDeptResp, 0)
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.id=sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, lang.SuccessCode, nil
}

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

// UpdateStatus 修改SysRole对象status
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

// GetWithName 获取SysRole对象
func (e *SysRole) GetWithName(d *dto.SysRoleQueryReq) (*models.SysRole, int, error) {
	model := &models.SysRole{}
	err := e.Orm.Where("role_name = ?", d.RoleName).First(model).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	menuIds, respCode, err := e.GetRoleMenuId(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	deptIds, respCode, err := e.GetRoleDeptId(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	model.MenuIds = menuIds
	model.DeptIds = deptIds
	return model, lang.SuccessCode, nil
}

// GetPermissionsById 获取权限对象
func (e *SysRole) GetPermissionsById(roleId int64) ([]string, int, error) {
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
