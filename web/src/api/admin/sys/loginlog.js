import request from '@/utils/request'

// 查询SysLoginlog列表
export function listSysLoginlog(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-login-log',
    method: 'get',
    params: query
  })
}

// 删除SysLoginlog
export function delSysLoginlog(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-login-log',
    method: 'delete',
    data: data
  })
}

export function exportSysLoginlog(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-login-log/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}
