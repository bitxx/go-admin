import request from '@/utils/request'

// 查询SysApi列表
export function listSysApi(query) {
  return request({
    url: '/admin-api/v1/sys-api',
    method: 'get',
    params: query
  })
}

// 查询SysApi详细
export function getSysApi(id) {
  return request({
    url: '/admin-api/v1/sys-api/' + id,
    method: 'get'
  })
}

// 修改SysApi
export function updateSysApi(data) {
  return request({
    url: '/admin-api/v1/sys-api/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除SysApi
export function delSysApi(data) {
  return request({
    url: '/admin-api/v1/sys-api',
    method: 'delete',
    data: data
  })
}

// 导出Announcement列表
export function exportSysApi(query) {
  return request({
    url: '/admin-api/v1/sys-api/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}
