import request from '@/utils/request'

// 查询参数列表
export function listConfig(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config',
    method: 'get',
    params: query
  })
}

// 查询参数详细
export function getConfig(configId) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config/' + configId,
    method: 'get'
  })
}

// 根据参数键名查询参数值
export function getConfigKey(configKey) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config/' + configKey,
    method: 'get'
  })
}

// 新增参数配置
export function addConfig(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config',
    method: 'post',
    data: data
  })
}

// 修改参数配置
export function updateConfig(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除参数配置
export function delConfig(data) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config',
    method: 'delete',
    data: data
  })
}

// 导出字典数据
export function exportSetConfig(query) {
  return request({
    url: '/admin-api/v1/admin/sys/sys-config/export',
    method: 'get',
    responseType: 'blob',
    params: query
  })
}
