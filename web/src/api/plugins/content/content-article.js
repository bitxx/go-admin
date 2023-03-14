import request from '@/utils/request'

// 查询ContentArticle列表
export function listContentArticle(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article',
    method: 'get',
    params: query
  })
}

// 查询ContentArticle详细
export function getContentArticle(id) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article/' + id,
    method: 'get'
  })
}

// 新增ContentArticle
export function addContentArticle(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article',
    method: 'post',
    data: data
  })
}

// 修改ContentArticle
export function updateContentArticle(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ContentArticle
export function delContentArticle(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article',
    method: 'delete',
    data: data
  })
}

// 导出ContentArticle列表
export function exportContentArticle(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-article/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

