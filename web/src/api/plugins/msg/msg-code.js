import request from '@/utils/request'

// 查询MsgCode列表
export function listMsgCode(query) {
  return request({
    url: '/admin-api/v1/plugins/msg/msg-code',
    method: 'get',
    params: query
  })
}
