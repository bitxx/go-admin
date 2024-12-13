import request from '@/utils/request'

// 查询菜单列表
export function listMenu(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu',
    method: 'get',
    params: query
  })
}

// 查询菜单详细
export function getMenu(menuId) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu/' + menuId,
    method: 'get'
  })
}

// 根据角色ID查询菜单下拉树结构
export function roleMenuTreeselect(roleId) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu/role-menu-tree-select/' + roleId,
    method: 'get'
  })
}

// 新增菜单
export function addMenu(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu',
    method: 'post',
    data: data
  })
}

// 修改菜单
export function updateMenu(data, id) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu/' + id,
    method: 'put',
    data: data
  })
}

// 删除菜单
export function delMenu(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu',
    method: 'delete',
    data: data
  })
}

export function getMenuRole() {
  return request({
    url: '/admin-api/v1/admin/sys/sys-menu/menu-role',
    method: 'get'
  })
}
