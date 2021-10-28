import {createRouter, createWebHistory} from 'vue-router'
import util from '../libs/util'

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        isAdmin?: boolean
        // 每个路由都必须声明
        requiresAuth: boolean
    }
}

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            meta: {
                requiresAuth: true,
            },
            component: () => import('../view/home.vue')
        },
        {
            path: '/app',
            name: 'app',
            meta: {
                requiresAuth: true,
            },
            component: () => import('../view/demo.vue')
        },
        {
            path: '/wx',
            name: 'wx',
            component: () => import('../view/wx.vue')
        },
        {
            path: '/login/:uuid?',
            name: 'login',
            component: () => import('../view/login.vue')
        },
        {
            path: '/register/:uuid?',
            name: 'register',
            component: () => import('../view/register.vue')
        },
        {
            path: '/:path(.*)',
            name: '404',
            component: () => import('../view/404.vue')
        }
        //...
    ],
})

router.beforeEach((to, from) => {
    // 而不是去检查每条路由记录
    // to.matched.some(record => record.meta.requiresAuth)
    if (to.meta.requiresAuth && !util.checkLogin()) {
        // 此路由需要授权，请检查是否已登录
        // 如果没有，则重定向到登录页面
        return {
            name: 'login',
            // 保存我们所在的位置，以便以后再来
            query: {redirect: to.fullPath},
        }
    }
})

export default router
