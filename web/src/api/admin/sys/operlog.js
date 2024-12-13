import request from '@/utils/request'

// 查询SysOperlog列表
export function listSysOperlog(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-oper-log',
    method: 'get',
    params: query
  })
}

// 删除SysOperlog
export function delSysOperlog(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-oper-log',
    method: 'delete',
    data: data
  })
}

export function exportSysOperlog(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-oper-log/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

