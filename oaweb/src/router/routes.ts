import { Auths } from 'src/models/auth';
import { RouteRecordRaw } from 'vue-router';
import menus from 'src/layouts/menus'

declare module 'vue-router' {
  interface RouteMeta {
    // 是可选的
    isAdmin?: boolean
    title?: string
    // 每个路由都必须声明
    requiresAuth: boolean
    checkAuth?: (a: Auths, r?: RouteLocationNormalized) => boolean
  }
}

function loadcomponents(path: string, name: string, main: string) {
  return {
    path: path,
    name: name,
    component: () => import("../pages/" + main + ".vue"),
  }
}
const routes: RouteRecordRaw[] = [

  {
    path: '/',
    component: () => import('src/layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
    },
    redirect: 'home',
    children: [
      {
        path: 'app/:id',
        name: 'app',
        component: () => import("../layouts/AppLayout.vue"),
        redirect: { name: 'app.home' },
        children: [
        ]
      },
      loadcomponents('home', 'home', 'IndexPage'),
      loadcomponents('user', 'user', 'user'),
      loadcomponents('fs', 'fs', 'fs'),
      loadcomponents('doc', 'doc', 'doc'),
      loadcomponents('stats', 'stats', 'stats'),
      loadcomponents('doc/:typ/:url(.*)', 'doc_item', 'docItem'),
      loadcomponents('settings', 'settings', 'settings'),
    ],
  },
  loadcomponents('/login/:uuid?', 'login', 'login'),
  loadcomponents('/register/:uuid?', 'register', 'register'),
  loadcomponents('/:catchAll(.*)*', '404', '404')
];
for (let i of menus.appLinks.value) {
  // @ts-ignore
  routes[0].children[0].children.push(i.router)
}
for (let i in menus.uniqueLinks) {
  for (let j of menus.uniqueLinks[i]) {
    // @ts-ignore
    routes[0].children[0].children.push(j.router)
  }
}

export default routes;
