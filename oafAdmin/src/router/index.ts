/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:14
* @description：index
*/


import { createRouter, createWebHistory } from 'vue-router'
import { Auths, R } from '@/auth'
import { useAppStore } from '@/store/app'
import util from '@/libs/util'
import { useUserStore } from '@/store/user'
import msg from '@/msg'


declare module 'vue-router' {
    interface RouteMeta {
        title?: string
        isAdmin?: boolean
        requiresAuth: boolean
        checkAuth?: (a: Auths, r?: RouteLocationNormalized) => boolean
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
            component: () => import('@/views/home.vue'),
        },
        {
            path: '/app/:uuid?',
            component: () => import('@/views/app.vue'),
            redirect: { name: 'app.main' },
            children: [
                {
                    path: 'main',
                    name: 'app.main',
                    meta: {
                        title: '首页',
                        requiresAuth: true,
                    },
                    component: () => import('@/views/app/main.vue'),
                },
                {
                    path: 'users',
                    name: 'app.users',
                    meta: {
                        title: '用户',
                        requiresAuth: true,
                        checkAuth: (a, r) => {
                            return a.Get(R.User, r?.params.uuid as string).CanRead()
                        },
                    },
                    component: () => import('@/views/app/users.vue'),
                },
                {
                    path: 'roles',
                    name: 'app.roles',
                    meta: {
                        title: '权限',
                        requiresAuth: true,
                        checkAuth: (a, r) => {
                            return a.Get(R.Role, r?.params.uuid as string).CanRead()
                        },
                    },
                    component: () => import('@/views/app/roles.vue'),
                },
                {
                    path: 'setting',
                    name: 'app.setting',
                    meta: {
                        title: '应用设置',
                        requiresAuth: true,
                        checkAuth: (a, r) => {
                            return a.Get(R.App, r?.params.uuid as string).CanUpdate()
                        },
                    },
                    component: () => import('@/views/app/setting.vue'),
                },
            ],
        },
        {
            path: '/user/setting',
            name: 'user_setting',
            meta: {
                requiresAuth: true,
            },
            component: () => import('@/views/user_setting.vue'),
        },
        {
            path: '/about',
            name: 'about',
            component: () => import('@/views/about.vue'),
        },
        {
            path: '/wx',
            name: 'wx',
            component: () => import('@/views/wx.vue'),
        },
        {
            path: '/login/:uuid?',
            name: 'login',
            component: () => import('@/views/login.vue'),
        },
        {
            path: '/register/:uuid?',
            name: 'register',
            component: () => import('@/views/register.vue'),
        },
        {
            path: '/:path(.*)',
            name: '404',
            component: () => import('@/views/404.vue'),
        },
    ],
})

let user: any = null
router.beforeEach((to, from) => {
    // to.matched.some(record => record.meta.requiresAuth)
    if (to.query.noh === '1') {
        let app = useAppStore()
        app.hideHeader = true
    }
    if (to.meta.requiresAuth && !util.checkLogin()) {
        // 此路由需要授权，请检查是否已登录
        // 如果没有，则重定向到登录页面
        return {
            name: 'login',
            // 保存我们所在的位置，以便以后再来
            query: { redirect: to.fullPath },
        }
    }
    if (to.meta.checkAuth) {
        if (!user) {
            user = useUserStore()
        }
        if (!to.meta.checkAuth(user.auth, to)) {

            msg.Warn('无权访问')
            return from
        }
    }
})

export default router
