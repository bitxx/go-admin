import request from '@/utils/request'

// 查询ContentAnnouncement列表
export function listContentAnnouncement(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement',
    method: 'get',
    params: query
  })
}

// 查询ContentAnnouncement详细
export function getContentAnnouncement(id) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement/' + id,
    method: 'get'
  })
}

// 新增ContentAnnouncement
export function addContentAnnouncement(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement',
    method: 'post',
    data: data
  })
}

// 修改ContentAnnouncement
export function updateContentAnnouncement(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ContentAnnouncement
export function delContentAnnouncement(data) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement',
    method: 'delete',
    data: data
  })
}

// 导出ContentAnnouncement列表
export function exportContentAnnouncement(query) {
  return request({
    url: '/admin-api/v1/plugins/content/content-announcement/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

