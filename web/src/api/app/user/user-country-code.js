import request from '@/utils/request'

// 查询UserCountryCode列表
export function listUserCountryCode(query) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code',
    method: 'get',
    params: query
  })
}

// 查询UserCountryCode详细
export function getUserCountryCode(id) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code/' + id,
    method: 'get'
  })
}

// 新增UserCountryCode
export function addUserCountryCode(data) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code',
    method: 'post',
    data: data
  })
}

// 修改UserCountryCode
export function updateUserCountryCode(data) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除UserCountryCode
export function delUserCountryCode(data) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code',
    method: 'delete',
    data: data
  })
}

// 导出UserCountryCode列表
export function exportUserCountryCode(query) {
  return request({
    url: '/admin-api/v1/app/user/user-country-code/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

