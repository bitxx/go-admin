import request from '@/utils/request'
import { data } from 'autoprefixer'

export const sysfiledirList = data => request({
  url: '/api-admin/v1/sysfiledir',
  method: 'GET',
  data
})

export const sysfiledirAcionAdd = data => request({
  url: '/api-admin/v1/sysfiledir',
  method: 'POST',
  data
})

export const sysfiledirAcionEdit = data => request({
  url: '/api-admin/v1/sysfiledir/' + data.id,
  method: 'PUT',
  data
})

export const sysfiledirAcionDel = data => request({
  url: '/api-admin/v1/sysfiledir/' + data,
  method: 'DELETE'
})

export const sysfileinfoList = data => request({
  url: `/api-admin/v1/sysfileinfo?pId=${data.pId}&pageSize=${data.pageSize}&pageIndex=${data.pageIndex}`,
  method: 'GET',
  data
})

export const sysfileinfo = id => request({
  url: '/api-admin/v1/sysfileinfo/' + id,
  method: 'GET'
})

export const sysfileinfoAdd = data => request({
  url: '/api-admin/v1/sysfileinfo',
  method: 'POST',
  data
})

export const sysfileinfoEdit = data => request({
  url: '/api-admin/v1/sysfileinfo/' + data.id,
  method: 'put',
  data
})

export const sysfileinfoDelete = id => request({
  url: '/api-admin/v1/sysfileinfo/' + id,
  method: 'DELETE',
  data
})
