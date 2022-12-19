/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:14
* @description：index
*/


import {createRouter, createWebHistory} from 'vue-router'

declare module 'vue-router' {
    interface RouteMeta {
        title?: string
        requiresAuth: boolean
        checkAuth?: (r?: RouteLocationNormalized) => boolean
    }
}
const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/home.vue'),
        },
        {
            path: '/:path(.*)',
            name: '404',
            component: () => import('@/views/404.vue'),
        },
    ],
})

router.beforeEach((to, from) => {
    if (to.meta.requiresAuth && to.meta.checkAuth) {
        return {
            name: 'login',
            // 保存我们所在的位置，以便以后再来
            query: {redirect: to.fullPath},
        }
    }
})

export default router
