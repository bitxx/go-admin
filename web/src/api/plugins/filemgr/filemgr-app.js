import request from '@/utils/request'

// 查询FilemgrApp列表
export function listFilemgrApp(query) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app',
    method: 'get',
    params: query
  })
}

// 查询FilemgrApp详细
export function getFilemgrApp(id) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app/' + id,
    method: 'get'
  })
}

// 新增FilemgrApp
export function addFilemgrApp(data) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app',
    method: 'post',
    data: data
  })
}

// 修改FilemgrApp
export function updateFilemgrApp(data) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除FilemgrApp
export function delFilemgrApp(data) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app',
    method: 'delete',
    data: data
  })
}

// 导出FilemgrApp列表
export function exportFilemgrApp(query) {
  return request({
    url: '/admin-api/v1/plugins/filemgr/filemgr-app/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}

