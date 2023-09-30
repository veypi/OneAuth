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
const routes: RouteRecordRaw[] = [

  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        component: () => import('pages/IndexPage.vue')
      }
    ],
  },
  {
    path: '/login/:uuid?',
    name: 'login',
    component: () => import('pages/login.vue'),
  },
  {
    path: '/register/:uuid?',
    name: 'register',
    component: () => import('pages/register.vue'),
  },


  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
