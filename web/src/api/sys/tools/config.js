import service from '@/utils/request'

// @Tags system
// @Summary 获取附u我
export const getConfig = () => {
  return service({
    url: '/v1/sysRuntimeConfig/getConfig',
    method: 'get',
    donNotShowLoading: true
  })
}
