import { route } from 'quasar/wrappers';
import util from 'src/libs/util';
import { useUserStore } from 'src/stores/user';
import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';

import routes from './routes';

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

function newRouter(/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : (process.env.VUE_ROUTER_MODE === 'history' ? createWebHistory : createWebHashHistory);

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });

  Router.beforeEach((to, from) => {
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
      const u = useUserStore()
      if (!to.meta.checkAuth(u.auth, to)) {

        // if (window.$msg) {
        //   window.$msg.warning('无权访问')
        // }
        return from
      }
    }
  })
  return Router;
};

export default newRouter();
