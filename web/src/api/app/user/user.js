import request from '@/utils/request'

// 查询User列表
export function listUser(query) {
  return request({
    url: '/admin-api/v1/app/user/user',
    method: 'get',
    params: query
  })
}

// 查询User详细
export function getUser(id) {
  return request({
    url: '/admin-api/v1/app/user/user/' + id,
    method: 'get'
  })
}

// 新增User
export function addUser(data) {
  return request({
    url: '/admin-api/v1/app/user/user',
    method: 'post',
    data: data
  })
}

// 修改User
export function updateUser(data) {
  return request({
    url: '/admin-api/v1/app/user/user/' + data.id,
    method: 'put',
    data: data
  })
}

// 导出User列表
export function exportUser(query) {
  return request({
    url: '/admin-api/v1/app/user/user/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

