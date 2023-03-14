import request from '@/utils/request'

// 查询UserConf列表
export function listUserConf(query) {
  return request({
    url: '/admin-api/v1/app/user/user-conf',
    method: 'get',
    params: query
  })
}

// 查询UserConf详细
export function getUserConf(id) {
  return request({
    url: '/admin-api/v1/app/user/user-conf/' + id,
    method: 'get'
  })
}

// 修改UserConf
export function updateUserConf(data) {
  return request({
    url: '/admin-api/v1/app/user/user-conf/' + data.id,
    method: 'put',
    data: data
  })
}

