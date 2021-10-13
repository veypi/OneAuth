import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Home from '../views/Home.vue'
import Demo from '@/views/demo.vue'
import Login from '@/views/login.vue'
import Register from '@/views/register.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/app',
    name: 'app',
    component: Demo
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/register',
    name: 'register',
    component: Register
  },
  {
    path: '/wx',
    name: 'wx',
    component: () => import('../views/wx.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
