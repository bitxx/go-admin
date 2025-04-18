package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	mycasbin "go-admin/core/casbin"
	"go-admin/core/dto/api"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
)

type SysMenu struct {
	api.Api
}

// GetTreeList admin-获取菜单管理树
func (e SysMenu) GetTreeList(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	list, respCode, err := s.GetTreeList(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get admin-获取菜单管理详情
func (e SysMenu) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert admin-新增菜单管理
func (e SysMenu) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update admin-更新菜单管理
func (e SysMenu) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	cb := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
	b, respCode, err := s.Update(&req, p, cb)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	_, _ = mycasbin.LoadPolicy(c)
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete admin-删除菜单管理
func (e SysMenu) Delete(c *gin.Context) {
	req := dto.SysMenuDeleteReq{}
	s := service.SysMenu{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Delete(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// GetMenuRole admin-根据角色获取菜单
func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	result, respCode, err := s.GetMenuRole(auth.Auth.GetRoleKey(c))

	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if result == nil || len(result) <= 0 {
		e.Error(baseLang.SysNoRoleMenuCode, lang.MsgErr(baseLang.SysNoRoleMenuCode, e.Lang).Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// GetMenuTreeSelect admin-获取全部菜单以及选中的菜单编号
func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectMenuRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	result, respCode, err := m.GetTreeList(&dto.SysMenuQueryReq{})
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}

	var menuIds []int64
	if req.RoleId > 0 {
		menuIds, respCode, err = r.GetMenuIdsByRole(req.RoleId)
		if err != nil {
			e.Error(respCode, err.Error())
			return
		}
	}

	e.OK(dto.MenuTreeRoleResp{
		Menus:       result,
		CheckedKeys: menuIds,
	}, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
