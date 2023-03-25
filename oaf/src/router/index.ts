import { createRouter, createWebHistory } from 'vue-router'
import util from '@/libs/util'
import { Auths, R } from '@/auth'
import { store } from '@/store'


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
                            return a.Get(R.User, r.params.uuid as string).CanRead()
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
                            return a.Get(R.Role, r.params.uuid as string).CanRead()
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
                            return a.Get(R.App, r.params.uuid as string).CanUpdate()
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

router.beforeEach((to, from) => {
    // to.matched.some(record => record.meta.requiresAuth)
    if (to.query.noh === '1') {
        store.commit('hideHeader', true)
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
        if (!to.meta.checkAuth(store.state.user.auth, to)) {

            // @ts-ignore
            if (window.$msg) {
                // @ts-ignore
                window.$msg.warning('无权访问')
            }
            return from
        }
    }
})

export default router
