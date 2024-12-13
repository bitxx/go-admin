import storage from '@/utils/storage'
const state = {
  info: storage.get('app_info')
}

const mutations = {
  SET_INFO: (state, data) => {
    state.info = data
    storage.set('app_info', data)
  }
}

const actions = {
  settingDetail({ commit }) {
    return new Promise((resolve, reject) => {
      var data = { admin_sys_app_name: 'go-admin后台管理系统', admin_sys_app_logo: 'http://www.bitxx.top/images/my_head-touch-icon-next.png' }
      commit('SET_INFO', data)
      resolve(data)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
