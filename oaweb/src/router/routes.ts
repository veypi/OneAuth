import { Auths } from 'src/models/auth';
import { RouteRecordRaw } from 'vue-router';

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
    components: {
      default: () => import("../pages/" + main + ".vue"),
    }
  }
}
const routes: RouteRecordRaw[] = [

  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
    },
    redirect: 'home',
    children: [
      loadcomponents('home', 'home', 'IndexPage'),
      loadcomponents('user', 'user', '404'),
      loadcomponents('settings', 'settings', '404'),
      {
        path: 'app/:id?',
        component: () => import("../layouts/AppLayout.vue"),
        redirect: { name: 'app.home' },
        children: [
          loadcomponents('home', 'app.home', 'IndexPage'),
          loadcomponents('user', 'app.user', 'AppHome'),
        ]
      }
    ],
  },
  {
    path: '/login/:uuid?',
    name: 'login',
    component: () => import('../pages/login.vue'),
  },
  {
    path: '/register/:uuid?',
    name: 'register',
    component: () => import('../pages/register.vue'),
  },


  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('../pages/404.vue'),
  },
];

export default routes;
