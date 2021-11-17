import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Home from '../views/Home.vue'
import Demo from '@/views/demo.vue'
import Login from '@/views/login.vue'
import Register from '@/views/register.vue'
import NotFound from '@/views/404.vue'

Vue.use(VueRouter)
// 避免push到相同路径报错
// 获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
// 修改原型对象中的push方法
VueRouter.prototype.push = function push(location: any) {
  // eslint-disable-next-line
  // @ts-ignore
  return originalPush.call(this, location).catch(err => err)
}

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/app',
    name: 'app',
    component: Demo
  },
  {
    path: '/login/:uuid?',
    name: 'login',
    component: Login
  },
  {
    path: '/register/:uuid?',
    name: 'register',
    component: Register
  },
  {
    path: '/wx',
    name: 'wx',
    component: () => import('../views/wx.vue')
  },
  {
    path: '*',
    name: '404',
    component: NotFound
  }
]

const router = new VueRouter({
  routes
})

export default router
