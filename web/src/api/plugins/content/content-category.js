import request from '@/utils/request'

// 查询ContentCategory列表
export function listContentCategory(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category',
    method: 'get',
    params: query
  })
}

// 查询ContentCategory详细
export function getContentCategory(id) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category/' + id,
    method: 'get'
  })
}

// 新增ContentCategory
export function addContentCategory(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category',
    method: 'post',
    data: data
  })
}

// 修改ContentCategory
export function updateContentCategory(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ContentCategory
export function delContentCategory(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category',
    method: 'delete',
    data: data
  })
}

// 导出ContentCategory列表
export function exportContentCategory(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-category/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

