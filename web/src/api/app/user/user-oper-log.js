import request from '@/utils/request'

// 查询UserOperLog列表
export function listUserOperLog(query) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log',
    method: 'get',
    params: query
  })
}

// 查询UserOperLog详细
export function getUserOperLog(id) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log/' + id,
    method: 'get'
  })
}

// 新增UserOperLog
export function addUserOperLog(data) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log',
    method: 'post',
    data: data
  })
}

// 修改UserOperLog
export function updateUserOperLog(data) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除UserOperLog
export function delUserOperLog(data) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log',
    method: 'delete',
    data: data
  })
}

// 导出UserOperLog列表
export function exportUserOperLog(query) {
  return request({
    url: '/admin-api/v1/app/user/user-oper-log/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

