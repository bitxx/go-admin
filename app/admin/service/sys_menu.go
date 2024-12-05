package service

import (
	"go-admin/app/admin/constant"
	sysLang "go-admin/app/admin/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"sort"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/core/dto"
)

type SysMenu struct {
	service.Service
}

func NewSysMenuService(s *service.Service) *SysMenu {
	var srv = new(SysMenu)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetTreeList 获取SysMenu列表 一次性获取所有数据
func (e *SysMenu) GetTreeList(c *dto.SysMenuQueryReq) ([]models.SysMenu, int, error) {
	list, respCode, err := e.GetList(c, false)
	if err != nil {
		return nil, respCode, err
	}

	count := len(list)
	var menus []models.SysMenu
	sort.Slice(list, func(i, j int) bool {
		return list[i].Sort < list[j].Sort
	})
	for i := 0; i < count; i++ {
		menusInfo := menuCall(&list, list[i])
		menus = append(menus, menusInfo)
	}
	//menus := menuCall2(list)
	return menus, lang.SuccessCode, nil
}

func menuCall2(menuList []models.SysMenu) []models.SysMenu {
	// 创建一个 map 来存储每个 parentId 对应的子菜单
	menuMap := make(map[int64][]models.SysMenu)
	var rootItems []models.SysMenu

	// 遍历菜单项，根据 parentId 分组
	for _, item := range menuList {
		menuMap[item.ParentId] = append(menuMap[item.ParentId], item)
	}

	// 遍历菜单项，递归构建子菜单
	for index, _ := range menuList {
		item := menuList[index]
		if item.ParentId <= 0 {
			// 如果 parentId == 0，则是根菜单项
			rootItems = append(rootItems, item)
		} else {
			// 为当前菜单项添加子菜单
			if children, exists := menuMap[item.Id]; exists {
				item.Children = children
			}
		}
	}

	return rootItems
}

// Get 获取SysMenu对象
func (e *SysMenu) Get(id int64, p *middleware.DataPermission) (*models.SysMenu, int, error) {
	if id <= 0 {
		//id<=0,表示为顶级根菜单
		return &models.SysMenu{Id: 0, ParentIds: ""}, lang.SuccessCode, nil
	}
	data := &models.SysMenu{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Preload("SysApi").First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	apis := make([]int, 0)
	for _, v := range data.SysApi {
		apis = append(apis, v.Id)
	}
	data.Apis = apis
	return data, lang.SuccessCode, nil
}

// QueryOne 通过自定义条件获取一条记录
func (e *SysMenu) QueryOne(queryCondition *dto.SysMenuQueryReq, p *middleware.DataPermission) (*models.SysMenu, int, error) {
	data := &models.SysMenu{}
	err := e.Orm.Model(&models.SysMenu{}).
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
func (e *SysMenu) Count(c *dto.SysMenuQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysMenu{}).
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

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(c *dto.SysMenuInsertReq) (int64, int, error) {
	if c.ParentId < 0 {
		return 0, sysLang.SysMenuParentIdEmptyCode, lang.MsgErr(sysLang.SysMenuParentIdEmptyCode, e.Lang)
	}
	if c.Title == "" {
		return 0, sysLang.SysMenuTitleEmptyCode, lang.MsgErr(sysLang.SysMenuTitleEmptyCode, e.Lang)
	}
	if c.MenuType == "" {
		return 0, sysLang.SysMenuTypeEmptyCode, lang.MsgErr(sysLang.SysMenuTypeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return 0, sysLang.SysMenuSortEmptyCode, lang.MsgErr(sysLang.SysMenuSortEmptyCode, e.Lang)
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
		data.Path = c.Path
		data.Element = c.Element
		data.IsHidden = c.IsHidden
		if c.MenuType == constant.MenuM {
			data.Redirect = c.Redirect
		}
		if c.MenuType == constant.MenuC {
			c.IsKeepAlive = global.SysStatusNotOk
			c.IsAffix = global.SysStatusNotOk
			data.IsKeepAlive = c.IsKeepAlive
			data.IsAffix = c.IsAffix
			data.Name = c.Name
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改SysMenu对象
func (e *SysMenu) Update(c *dto.SysMenuUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Title == "" {
		return false, sysLang.SysMenuTitleEmptyCode, lang.MsgErr(sysLang.SysMenuTitleEmptyCode, e.Lang)
	}
	if c.MenuType == "" {
		return false, sysLang.SysMenuTypeEmptyCode, lang.MsgErr(sysLang.SysMenuTypeEmptyCode, e.Lang)
	}
	if c.Sort < 0 {
		return false, sysLang.SysMenuSortEmptyCode, lang.MsgErr(sysLang.SysMenuSortEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var alist = make([]models.SysApi, 0)
	tx.Where("id in ?", c.Apis).Find(&alist)

	err = tx.Model(&data).Association("SysApi").Delete(data.SysApi)
	if err != nil {
		return false, lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}

	m, respCode, err := e.Get(c.ParentId, p)
	if err != nil {
		return false, respCode, err
	}

	now := time.Now()
	if c.MenuType == constant.MenuM || c.MenuType == constant.MenuC {
		data.Path = c.Path
		data.Element = c.Element
		data.IsHidden = c.IsHidden
		if c.MenuType == constant.MenuM {
			data.Redirect = c.Redirect
		}
		if c.MenuType == constant.MenuC {
			c.IsKeepAlive = global.SysStatusNotOk
			c.IsAffix = global.SysStatusNotOk
			data.IsKeepAlive = c.IsKeepAlive
			data.IsAffix = c.IsAffix
			data.Name = c.Name
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
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	err = tx.Model(&data).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&data).Error
	if err != nil {
		return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
	}
	return true, lang.SuccessCode, nil
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
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
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return sysLang.SysMenuHasChildCode, lang.MsgErr(sysLang.SysMenuHasChildCode, e.Lang)
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
			return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
		}
	}

	//删除列表
	var data models.SysMenu
	err = tx.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetList 获取菜单数据
func (e *SysMenu) GetList(c *dto.SysMenuQueryReq, withApi bool) ([]models.SysMenu, int, error) {
	var list []models.SysMenu
	var err error
	if withApi {
		err = e.Orm.Model(&models.SysMenu{}).Preload("SysApi").
			Scopes(
				cDto.MakeCondition(c.GetNeedSearch()),
			).Find(&list).Error
	} else {
		err = e.Orm.Model(&models.SysMenu{}).
			Scopes(
				cDto.MakeCondition(c.GetNeedSearch()),
			).Find(&list).Error
	}
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}

// GetMenuLabelTree 获取菜单的完整树结构(用来显示简单的菜单信息：编号 名称)
// 角色添加或者更新时，选择菜单列表会用到，菜单权限
func (e *SysMenu) GetMenuLabelTree() ([]dto.MenuLabel, int, error) {
	var list []models.SysMenu
	list, respCode, err := e.GetList(&dto.SysMenuQueryReq{}, false)
	if err != nil {
		return nil, respCode, err
	}

	m := make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].Id
		e.Label = list[i].Title
		deptsInfo := menuLabelCall(&list, e)
		m = append(m, deptsInfo)
	}
	return m, lang.SuccessCode, nil
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList
	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].Id
		mi.Label = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != constant.MenuF {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

// menuCall 菜单完整数据-构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.Id = list[j].Id
		mi.Name = list[j].Name
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.ParentIds = list[j].ParentIds
		mi.IsFrame = list[j].IsFrame
		mi.IsKeepAlive = list[j].IsKeepAlive
		mi.IsAffix = list[j].IsAffix
		mi.Element = list[j].Element
		mi.Redirect = list[j].Redirect
		mi.Sort = list[j].Sort
		mi.IsHidden = list[j].IsHidden
		mi.CreatedAt = list[j].CreatedAt
		mi.SysApi = list[j].SysApi
		mi.Children = []models.SysMenu{}
		ms := menuCall(menuList, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

// SetMenuRole 获取左侧菜单树使用，后台主页管理菜单
func (e *SysMenu) SetMenuRole(roleKey string) ([]models.SysMenu, int, error) {
	menus, respCode, err := e.getByRoleKey(roleKey)
	m := make([]models.SysMenu, 0)
	for i := 0; i < len(menus); i++ {
		//使用parentId=0的开始铸造
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return m, respCode, err
}

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
			//menuList = *role.SysMenu

			temp := map[int64]int{}
			for _, v := range *role.SysMenu {
				ids := strings.Split(v.ParentIds, ",")
				for _, idStr := range ids {
					id, _ := strconv.ParseInt(idStr, 10, 64)
					temp[id] = temp[id] + 1
					if temp[id] == 1 && id > 0 {
						data := models.SysMenu{}
						err = e.Orm.Where("id=?", id).Find(&data).Error
						if data.MenuType == constant.MenuM || data.MenuType == constant.MenuC {
							menuList = append(menuList, data)
						}

					}
				}
			}
		}
	}

	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return menuList, lang.SuccessCode, nil
}
