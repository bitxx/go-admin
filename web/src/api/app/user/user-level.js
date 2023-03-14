import request from '@/utils/request'

// 查询UserLevel列表
export function listUserLevel(query) {
  return request({
    url: '/admin-api/v1/app/user/user-level',
    method: 'get',
    params: query
  })
}

// 查询UserLevel详细
export function getUserLevel(id) {
  return request({
    url: '/admin-api/v1/app/user/user-level/' + id,
    method: 'get'
  })
}

// 新增UserLevel
export function addUserLevel(data) {
  return request({
    url: '/admin-api/v1/app/user/user-level',
    method: 'post',
    data: data
  })
}

// 修改UserLevel
export function updateUserLevel(data) {
  return request({
    url: '/admin-api/v1/app/user/user-level/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除UserLevel
export function delUserLevel(data) {
  return request({
    url: '/admin-api/v1/app/user/user-level',
    method: 'delete',
    data: data
  })
}

// 导出UserLevel列表
export function exportUserLevel(query) {
  return request({
    url: '/admin-api/v1/app/user/user-level/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

