import { asyncRoutes, constantRoutes } from '@/router'
import Layout from '@/layout'
import { getMenuRole } from '@/api/sys/menu'
// import sysuserindex from '@/views/sysuser/index'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Use names to determine if the current user has permission
 * @param names
 * @param route
 */
function hasPathPermission(paths, route) {
  if (route.path) {
    return paths.some(path => route.path === path.path)
  } else {
    return true
  }
}

/**
  * 后台查询的菜单数据拼装成路由格式的数据
  * @param routes
  */
export function generaMenu(routes, data) {
  data.forEach(item => {
    var hidden = true
    var keepAlive = false
    if (item.isHidden === '2') {
      hidden = false
    }
    if (item.isKeepAlive === '1') {
      keepAlive = true
    }
    const menu = {
      path: item.path,
      component: item.element === 'Layout' ? Layout : loadView(item.element),
      hidden: hidden,
      children: [],
      name: item.path,
      meta: {
        title: item.title,
        icon: item.icon,
        keepAlive: keepAlive
      }
    }
    if (item.children) {
      generaMenu(menu.children, item.children)
    }
    routes.push(menu)
  })
}

export const loadView = (view) => { // 路由懒加载
  return (resolve) => require([`@/views${view}`], resolve)
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param components
 */
export function filterAsyncPathRoutes(routes, paths) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPathPermission(paths, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncPathRoutes(tmp.children, paths)
      }
      res.push(tmp)
    }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: [],
  defaultRoutes: [],
  topbarRouters: [],
  sidebarRouters: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  },
  SET_DEFAULT_ROUTES: (state, routes) => {
    state.defaultRoutes = constantRoutes.concat(routes)
  },
  SET_TOPBAR_ROUTES: (state, routes) => {
    // 顶部导航菜单默认添加统计报表栏指向首页
    // const index = [{
    //   path: 'dashboard',
    //   meta: { title: '统计报表', icon: 'dashboard' }
    // }]
    state.topbarRouters = routes // .concat(index)
  },
  SET_SIDEBAR_ROUTERS: (state, routes) => {
    state.sidebarRouters = routes
  }
}

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise(resolve => {
      let loadMenuData = []

      getMenuRole().then(response => {
        let data = response
        if (response.code !== 200) {
          this.$message({
            message: '菜单数据加载异常',
            type: 0
          })
        } else {
          data = response.data
          const loadMenuDataTmp = []
          Object.assign(loadMenuDataTmp, data)
          loadMenuData = loadMenuDataTmp.filter(item => item.id !== 119 && item.id !== 120)
          generaMenu(asyncRoutes, loadMenuData)
          asyncRoutes.push({ path: '*', redirect: '/', hidden: true })
          commit('SET_ROUTES', asyncRoutes)
          const sidebarRoutes = []
          generaMenu(sidebarRoutes, loadMenuData)
          commit('SET_SIDEBAR_ROUTERS', constantRoutes.concat(sidebarRoutes))
          commit('SET_DEFAULT_ROUTES', sidebarRoutes)
          commit('SET_TOPBAR_ROUTES', sidebarRoutes)
          resolve(asyncRoutes)
        }
      }).catch(error => {
        console.log(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
