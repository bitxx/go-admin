import request from '@/utils/request'

// 查询生成表数据
export function listTable(query) {
  return request({
    url: '/admin-api/v1/sys/table',
    method: 'get',
    params: query
  })
}
// 查询db数据库列表
export function listDbTable(query) {
  return request({
    url: '/admin-api/v1/sys/table/dbtables',
    method: 'get',
    params: query
  })
}

// 查询表详细信息
export function getGenTable(tableId) {
  return request({
    url: '/admin-api/v1/sys/table/' + tableId,
    method: 'get'
  })
}

// 修改代码生成信息
export function updateGenTable(data, id) {
  return request({
    url: '/admin-api/v1/sys/table/' + id,
    method: 'put',
    data: data
  })
}

// 导入表
export function importTable(data) {
  return request({
    url: '/admin-api/v1/sys/table',
    method: 'post',
    data: data
  })
}
// 预览生成代码
export function previewTable(tableId) {
  return request({
    url: '/admin-api/v1/sys/table/preview/' + tableId,
    method: 'get'
  })
}

// 生成代码
export function genCode(id) {
  return request({
    url: '/admin-api/v1/sys/table/gen/' + id,
    method: 'get'
  })
}

export function downloadCode(id) {
  return request({
    url: '/admin-api/v1/sys/table/gen/download/' + id,
    method: 'get',
    responseType: 'blob'
  })
}

// 删除表数据
export function delTable(data) {
  return request({
    url: '/admin-api/v1/sys/table',
    method: 'delete',
    data: data
  })
}

// 生成菜单到数据库
export function genDB(tableId) {
  return request({
    url: '/admin-api/v1/sys/table/gen/db/' + tableId,
    method: 'get'
  })
}
