import request from '@/utils/request'

// 查询UserAccountLog列表
export function listUserAccountLog(query) {
  return request({
    url: '/admin-api/v1/app/user/user-account-log',
    method: 'get',
    params: query
  })
}

// 导出UserAccountLog列表
export function exportUserAccountLog(query) {
  return request({
    url: '/admin-api/v1/app/user/user-account-log/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

