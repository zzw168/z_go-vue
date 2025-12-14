import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/utils/bus.js'
import { asyncMenu } from '@/api/menu'
import { defineStore } from 'pinia'
import { ref, watchEffect } from 'vue'
import pathInfo from '@/pathInfo.json'
import {useRoute} from "vue-router";
import {config} from "@/core/config.js";

const notLayoutRouterArr = []   // 不使用 layout 的菜单（如登录页、首页等）
const keepAliveRoutersArr = []  // 需要 keep-alive 的路由组件名
const nameMap = {}              //子菜单 keep-alive 时，记录父菜单也要 keep-alive

// 格式化后端菜单 → 路由节点
const formatRouter = (routes, routeMap, parent) => {
  // 为每个后端菜单补充字段
  routes &&
    routes.forEach((item) => {
      item.parent = parent
      item.meta.btns = item.btns
      item.meta.hidden = item.hidden
      // 若 defaultMenu=true 且为顶级菜单，则直接作为独立页面。
      if (item.meta.defaultMenu === true) {
        if (!parent) {
          item = { ...item, path: `/${item.path}` }
          notLayoutRouterArr.push(item)
        }
      }
      // 建立 name → route 映射表,方便后续 keep-alive 和 tab 功能使用。
      routeMap[item.name] = item
      // 递归处理 children
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap, item)
      }
    })
}

// 生成需要缓存的组件列表
const KeepAliveFilter = (routes) => {
  routes &&
    routes.forEach((item) => {
      // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
      if (
        (item.children && item.children.some((ch) => ch.meta.keepAlive)) ||
        item.meta.keepAlive
      ) {
        const path = item.meta.path
        keepAliveRoutersArr.push(pathInfo[path])
        nameMap[item.name] = pathInfo[path]
      }
      if (item.children && item.children.length > 0) {
        KeepAliveFilter(item.children)
      }
    })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([])
  const asyncRouterFlag = ref(0)
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    
    // 1. 首先添加原有的keepAlive配置
    keepArrTemp.push(...keepAliveRoutersArr)
    if (config.KeepAliveTabs) {
      // 当用户打开多个标签（Tab），需要缓存这些页面。
      history.forEach((item) => {
        // 2. 为所有history中的路由强制启用keep-alive
        // 通过routeMap获取路由信息，然后通过pathInfo获取组件名
        const routeInfo = routeMap[item.name]
        if (routeInfo && routeInfo.meta && routeInfo.meta.path) {
          const componentName = pathInfo[routeInfo.meta.path]
          if (componentName) {
            keepArrTemp.push(componentName)
          }
        }
        
        // 3. 如果子路由在tabs中打开，父路由也需要keepAlive
        if (nameMap[item.name]) {
          keepArrTemp.push(nameMap[item.name])
        }
      })
    }
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }


  const route = useRoute()

  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])

  const topMenu = ref([])

  const leftMenu = ref([])

  const menuMap = {}

  const topActive = ref('')

  const setLeftMenu = (name) => {
    sessionStorage.setItem('topActive', name)
    topActive.value = name
    leftMenu.value = []
    if (menuMap[name]?.children) {
      leftMenu.value = menuMap[name].children
    }
    return menuMap[name]?.children
  }

  // 根据当前路由，自动找到对应 topMenu
  const findTopActive = (menuMap, routeName) => {
    for (let topName in menuMap) {
      const topItem = menuMap[topName];
      if (topItem.children?.some(item => item.name === routeName)) {
        return topName;
      }
      const foundName = findTopActive(topItem.children || {}, routeName);
      if (foundName) {
        return topName;
      }
    }
    return null;
  };

  // 监听路由变化，重建顶部菜单（一级菜单）
  watchEffect(() => {
    let topActive = sessionStorage.getItem('topActive')
    // 初始化菜单内容，防止重复添加
    topMenu.value = [];
    asyncRouters.value[0]?.children.forEach((item) => {
      if (item.hidden) return
      menuMap[item.name] = item
      topMenu.value.push({ ...item, children: [] })
    })
    if (!topActive || topActive === 'undefined' || topActive === 'null') {
      topActive = findTopActive(menuMap, route.name);
    }
    setLeftMenu(topActive)
  })

  const routeMap = {}
  // 从后台获取动态路由
  const SetAsyncRouter = async () => {
    asyncRouterFlag.value++
    // layout 是基础框架（sidebar/header/main）
    const baseRouter = [
      {
        path: '/layout',
        name: 'layout',
        component: 'view/layout/index.vue',
        meta: {
          title: '底层layout'
        },
        children: []
      }
    ]
    // 从后端获取菜单
    const asyncRouterRes = await asyncMenu()
    const asyncRouter = asyncRouterRes.data.menus
    asyncRouter &&
      asyncRouter.push({
        path: 'reload',
        name: 'Reload',
        hidden: true,
        meta: {
          title: '',
          closeTab: true
        },
        component: 'view/error/reload.vue'
      })
    // 格式化菜单树. 构建层级结构、keepAlive、路由映射
    formatRouter(asyncRouter, routeMap)
    // 注入到 baseRouter
    baseRouter[0].children = asyncRouter
    // 不走 layout 的路由追加
    if (notLayoutRouterArr.length !== 0) {
      baseRouter.push(...notLayoutRouterArr)
    }
    // 动态注入到 vue-router（最关键一步）
    // 把 component 字符串变成真正的异步组件
    // 动态调用 router.addRoute()
    asyncRouterHandle(baseRouter)
    KeepAliveFilter(asyncRouter)
    asyncRouters.value = baseRouter
    return true
  }

  return {
    topActive,
    setLeftMenu,
    topMenu,
    leftMenu,
    asyncRouters,
    keepAliveRouters,
    asyncRouterFlag,
    SetAsyncRouter,
    routeMap
  }
})
