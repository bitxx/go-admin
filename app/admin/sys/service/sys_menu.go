package service

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/tree"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysMenu struct {
	service.Service
}

// NewSysMenuService admin-实例化菜单管理
func NewSysMenuService(s *service.Service) *SysMenu {
	var srv = new(SysMenu)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTreeList admin-获取菜单管理树
func (e *SysMenu) GetTreeList(c *dto.SysMenuQueryReq) ([]*models.SysMenu, int, error) {
	list, respCode, err := e.GetList(c, false)
	if err != nil {
		return nil, respCode, err
	}

	treeList := tree.GenTree(&list,
		func(item models.SysMenu) int64 { return item.Id },
		func(item models.SysMenu) int64 { return item.ParentId },
		func(item *models.SysMenu, children []*models.SysMenu) { item.Children = children },
	)

	return []*models.SysMenu{
		{Id: 0, Title: "主目录", ParentId: 0, Children: treeList},
	}, baseLang.SuccessCode, nil
}

// Get admin-获取菜单管理详情
func (e *SysMenu) Get(id int64, p *middleware.DataPermission) (*models.SysMenu, int, error) {
	if id <= 0 {
		//id<=0,表示为顶级根菜单
		return &models.SysMenu{Id: 0, ParentIds: ""}, baseLang.SuccessCode, nil
	}
	data := &models.SysMenu{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Preload("SysApi").First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	apis := make([]int64, 0)
	for _, v := range data.SysApi {
		apis = append(apis, v.Id)
	}
	data.Apis = apis
	return data, baseLang.SuccessCode, nil
}

// GetWithRoles admin-根据菜单获取对应的所有角色
func (e *SysMenu) GetWithRoles(id int64) (*models.SysMenu, int, error) {
	if id <= 0 {
		//id<=0,表示为顶级根菜单
		return &models.SysMenu{Id: 0, ParentIds: ""}, baseLang.SuccessCode, nil
	}
	data := &models.SysMenu{}
	err := e.Orm.Preload("SysRole").First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne admin-获取菜单管理一条记录
func (e *SysMenu) QueryOne(queryCondition *dto.SysMenuQueryReq, p *middleware.DataPermission) (*models.SysMenu, int, error) {
	data := &models.SysMenu{}
	err := e.Orm.Model(&models.SysMenu{}).
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

// Count admin-获取菜单管理数据总数
func (e *SysMenu) Count(c *dto.SysMenuQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysMenu{}).
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

// Insert admin-新增菜单管理
func (e *SysMenu) Insert(c *dto.SysMenuInsertReq) (int64, int, error) {
	if c.ParentId < 0 {
		return 0, baseLang.SysMenuParentIdEmptyCode, lang.MsgErr(baseLang.SysMenuParentIdEmptyCode, e.Lang)
	}
	if c.Title == "" {
		return 0, baseLang.SysMenuTitleEmptyCode, lang.MsgErr(baseLang.SysMenuTitleEmptyCode, e.Lang)
	}
	if c.MenuType == "" {
		return 0, baseLang.SysMenuTypeEmptyCode, lang.MsgErr(baseLang.SysMenuTypeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return 0, baseLang.SysMenuSortEmptyCode, lang.MsgErr(baseLang.SysMenuSortEmptyCode, e.Lang)
	}

	m, respCode, err := e.Get(c.ParentId, nil)
	if err != nil {
		return 0, respCode, err
	}

	var alist = make([]models.SysApi, 0)
	e.Orm.Where("id in ?", c.Apis).Find(&alist)

	now := time.Now()
	data := models.SysMenu{}
	if c.MenuType == constant.MenuM || c.MenuType == constant.MenuC {
		//确保路由地址不重复
		if c.Path != "" {
			query := dto.SysMenuQueryReq{}
			query.Path = c.Path
			count, respCode, err := e.Count(&query)
			if err != nil && respCode != baseLang.DataNotFoundCode {
				return 0, respCode, err
			}
			if count > 0 {
				return 0, baseLang.SysMenuPathExistCode, lang.MsgErr(baseLang.SysMenuPathExistCode, e.Lang)
			}
		}
		data.Path = c.Path
		data.IsHidden = c.IsHidden
		if c.MenuType == constant.MenuM {
			data.Redirect = c.Redirect
		}
		if c.MenuType == constant.MenuC {
			data.IsKeepAlive = c.IsKeepAlive
			data.Element = c.Element
			data.IsAffix = c.IsAffix
			data.IsFrame = c.IsFrame
		}
	}
	if c.MenuType == constant.MenuC || c.MenuType == constant.MenuF {
		data.SysApi = alist
		if c.MenuType == constant.MenuF {
			data.Permission = c.Permission
		}
	}
	data.Title = c.Title
	data.Icon = c.Icon
	data.MenuType = c.MenuType
	data.ParentId = c.ParentId
	data.ParentIds = m.ParentIds + strconv.FormatInt(m.Id, 10) + ","
	data.Sort = c.Sort
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

// Update admin-更新菜单管理
func (e *SysMenu) Update(c *dto.SysMenuUpdateReq, p *middleware.DataPermission, cb *casbin.SyncedEnforcer) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Title == "" {
		return false, baseLang.SysMenuTitleEmptyCode, lang.MsgErr(baseLang.SysMenuTitleEmptyCode, e.Lang)
	}
	if c.MenuType == "" {
		return false, baseLang.SysMenuTypeEmptyCode, lang.MsgErr(baseLang.SysMenuTypeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return false, baseLang.SysMenuSortEmptyCode, lang.MsgErr(baseLang.SysMenuSortEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	oldCids := data.ParentIds + strconv.FormatInt(c.Id, 10) + ","

	//获取上级菜单
	m, respCode, err := e.Get(c.ParentId, p)
	if err != nil {
		return false, respCode, err
	}

	//获接口编号对应的接口详情
	var alist = make([]models.SysApi, 0)
	if err = e.Orm.Where("id in ?", c.Apis).Find(&alist).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	menuUpdates := map[string]interface{}{}
	now := time.Now()
	if c.MenuType == constant.MenuM || c.MenuType == constant.MenuC {
		//确保路由地址唯一
		if c.Path != "" && data.Path != c.Path {
			req := dto.SysMenuQueryReq{}
			req.Path = c.Path
			resp, respCode, err := e.QueryOne(&req, p)
			if err != nil && respCode != baseLang.DataNotFoundCode {
				return false, respCode, err
			}
			if respCode == baseLang.SuccessCode && resp.Id != data.Id {
				return false, baseLang.SysMenuPathExistCode, lang.MsgErr(baseLang.SysMenuPathExistCode, e.Lang)
			}
		}
		if c.Path != "" && data.Path != c.Path {
			menuUpdates["path"] = c.Path
		}
		if c.IsHidden != "" && data.IsHidden != c.IsHidden {
			menuUpdates["is_hidden"] = c.IsHidden
		}

		if c.MenuType == constant.MenuM {
			if c.Redirect != "" && data.Redirect != c.Redirect {
				menuUpdates["redirect"] = c.Redirect
			}
		}
		if c.MenuType == constant.MenuC {
			if c.Element != "" && data.Element != c.Element {
				menuUpdates["element"] = c.Element
			}
			if c.IsKeepAlive != "" && data.IsKeepAlive != c.IsKeepAlive {
				menuUpdates["is_keep_alive"] = c.IsKeepAlive
			}
			if c.IsAffix != "" && data.IsAffix != c.IsAffix {
				menuUpdates["is_affix"] = c.IsAffix
			}
			if c.IsFrame != "" && data.IsFrame != c.IsFrame {
				menuUpdates["is_frame"] = c.IsFrame
			}
		}
	}
	if c.Title != "" && data.Title != c.Title {
		menuUpdates["title"] = c.Title
	}
	if c.Icon != "" && data.Icon != c.Icon {
		menuUpdates["icon"] = c.Icon
	}
	if c.MenuType != "" && data.MenuType != c.MenuType {
		menuUpdates["menu_type"] = c.MenuType
	}
	if c.ParentId >= 0 && data.ParentId != c.ParentId {
		menuUpdates["parent_id"] = c.ParentId
	}
	newPids := m.ParentIds + strconv.FormatInt(m.Id, 10) + ","
	if data.ParentIds != newPids {
		menuUpdates["parent_ids"] = newPids
	}
	if c.Sort >= 0 && data.Sort != c.Sort {
		menuUpdates["sort"] = c.Sort
	}

	if c.MenuType == constant.MenuC || c.MenuType == constant.MenuF {
		if c.MenuType == constant.MenuF {
			if c.Permission != "" && data.Permission != c.Permission {
				menuUpdates["permission"] = c.Permission
			}
		}
	}
	needUpdate := false
	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		if len(menuUpdates) > 0 {
			data.UpdateBy = c.CurrUserId
			data.UpdatedAt = &now
			needUpdate = true
			if err = tx.Model(&data).Where("id=?", data.Id).Updates(menuUpdates).Error; err != nil {
				return err
			}
		}
		if c.MenuType == constant.MenuC || c.MenuType == constant.MenuF {
			needUpdateApi := len(data.SysApi) != len(alist)
			if !needUpdateApi {
				apiMap := make(map[int64]bool)
				for _, api := range data.SysApi {
					apiMap[api.Id] = true
				}
				// 遍历 alist，检查每个 SysApi 是否在 data.SysApi 中存在
				for _, api := range alist {
					if _, exists := apiMap[api.Id]; !exists {
						needUpdateApi = true
						break
					}
				}
			}
			if needUpdateApi {
				needUpdate = true
				data.SysApi = alist
				// 清空旧的关联关系并添加新的关联关系
				if err = tx.Model(&data).Association("SysApi").Replace(data.SysApi); err != nil {
					return err
				}
				//接口数据需要更新，则casbin也需要更新
				if _, err = e.updateCasbinByMenu(c.Id, tx, cb); err != nil {
					return err
				}
			}
		}

		//其余所有包含cidsOld的菜单或者按钮，均替换为cidsNew。菜单上级变了，那对应的子节点的parentIds也得变
		newCids := data.ParentIds + strconv.FormatInt(c.Id, 10) + ","
		if oldCids != newCids {
			needUpdate = true
			if err = tx.Model(&models.SysMenu{}).
				Where("parent_ids LIKE ?", oldCids+"%").
				Update("parent_ids", gorm.Expr("REPLACE(parent_ids, ?, ?)", oldCids, newCids)).Error; err != nil {
			}
		}

		return nil
	})
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	if needUpdate {
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// UpdateCasbinByMenu menu的接口变动，则对应更新Casbin
func (e *SysMenu) updateCasbinByMenu(menuId int64, tx *gorm.DB, cb *casbin.SyncedEnforcer) (int, error) {
	data, respCode, err := e.GetWithRoles(menuId)
	if err != nil {
		return respCode, err
	}
	if len(data.SysRole) <= 0 {
		return baseLang.SuccessCode, nil
	}

	for _, role := range data.SysRole {
		//根据角色获取到对应的菜单
		roleService := NewSysRoleService(&e.Service)
		menuIds, respCode, err := roleService.GetMenuIdsByRole(role.Id)
		if err != nil {
			return respCode, err
		}

		//根据菜单获取每个菜单对应的api，多对多
		var mlist = make([]models.SysMenu, 0)
		if err = tx.Preload("SysApi").Where("id in ?", menuIds).Find(&mlist).Error; err != nil {
			return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return roleService.UpdateCasbin(mlist, role.RoleKey, tx, cb)
	}
	return baseLang.SuccessCode, nil

}

// Delete admin-删除菜单管理
func (e *SysMenu) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	//检测是否可删除
	req := dto.SysMenuQueryReq{}
	req.ParentIds = ids
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return baseLang.SysMenuHasChildCode, lang.MsgErr(baseLang.SysMenuHasChildCode, e.Lang)
	}

	//删除关联的api
	for _, id := range ids {
		var er error
		temp, respCode, er := e.Get(id, p)
		if er != nil {
			err = er
			return respCode, er
		}
		err = tx.Model(&temp).Association("SysApi").Delete(temp.SysApi)
		if err != nil {
			return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
		}
	}

	//删除列表
	var data models.SysMenu
	err = tx.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// GetList admin-获取菜单管理全部列表
func (e *SysMenu) GetList(c *dto.SysMenuQueryReq, withApi bool) ([]models.SysMenu, int, error) {
	var list []models.SysMenu
	var err error
	if withApi {
		err = e.Orm.Model(&models.SysMenu{}).Order("sort").Preload("SysApi").
			Scopes(
				cDto.MakeCondition(c.GetNeedSearch()),
			).Find(&list).Error
	} else {
		err = e.Orm.Model(&models.SysMenu{}).Order("sort").
			Scopes(
				cDto.MakeCondition(c.GetNeedSearch()),
			).Find(&list).Error
	}
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, baseLang.SuccessCode, nil
}

// GetMenuRole admin-根据角色获取菜单树使用
func (e *SysMenu) GetMenuRole(roleKey string) ([]*models.SysMenu, int, error) {
	menus, respCode, err := e.getByRoleKey(roleKey)
	if err != nil {
		return nil, respCode, err
	}
	return tree.GenTree(&menus,
		func(item models.SysMenu) int64 { return item.Id },
		func(item models.SysMenu) int64 { return item.ParentId },
		func(item *models.SysMenu, children []*models.SysMenu) { item.Children = children },
	), respCode, err
}

// getByRoleKey admin-内部方法，根据角色获取菜单树使用
func (e *SysMenu) getByRoleKey(roleKey string) ([]models.SysMenu, int, error) {
	var menuList []models.SysMenu
	var err error
	if roleKey == constant.RoleKeyAdmin {
		var data []models.SysMenu
		err = e.Orm.Where(" menu_type in (?)", []string{constant.MenuM, constant.MenuC}).Order("sort").Find(&data).Error
		menuList = data
	} else {
		var role models.SysRole
		role.RoleKey = roleKey
		err = e.Orm.Debug().Model(&role).Where("role_key = ? ", roleKey).Preload("SysMenu", func(db *gorm.DB) *gorm.DB {
			return db.Where(" menu_type in (?)", []string{constant.MenuM, constant.MenuC, constant.MenuF}).Order("sort")
		}).Find(&role).Error
		if role.SysMenu != nil {
			filterParentMenuIds := make(map[int64]bool) // 存储所有的父菜单 ID
			menuSet := make(map[int64]struct{})         // 用于快速判断 menuList 中的 ID 是否存在

			// 遍历角色的菜单
			for _, v := range *role.SysMenu {
				// 添加到 menuSet 以快速查重
				menuSet[v.Id] = struct{}{}
				menuList = append(menuList, v)

				// 分割 ParentIds 并处理
				for _, idStr := range strings.Split(v.ParentIds, ",") {
					id, err := strconv.ParseInt(idStr, 10, 64)
					if err != nil || id == 0 {
						continue // 忽略解析失败或根节点
					}
					if _, exists := menuSet[id]; !exists {
						filterParentMenuIds[id] = true
					}
				}
			}

			// 收集需要二次获取的 parent IDs
			parentIds := make([]int64, 0, len(filterParentMenuIds))
			for id := range filterParentMenuIds {
				if _, exists := menuSet[id]; !exists {
					parentIds = append(parentIds, id)
				}
			}

			if len(parentIds) > 0 {
				var parentMenus []models.SysMenu
				menuTypes := []string{constant.MenuM, constant.MenuC}
				err = e.Orm.Where("id in (?) and menu_type in (?)", parentIds, menuTypes).Find(&parentMenus).Error
				if err == nil {
					menuList = append(menuList, parentMenus...)
				}
			}

		}
	}

	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return menuList, baseLang.SuccessCode, nil
}
