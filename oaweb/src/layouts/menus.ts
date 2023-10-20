/*
 * menu.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-18 22:58
 * Distributed under terms of the MIT license.
 */

import { Dict } from 'src/models';
import { ref } from 'vue';
import { RouteLocationRaw } from 'vue-router';
import cfg from 'src/cfg'

export interface MenuLink {
  title: string;
  caption?: string;
  to?: RouteLocationRaw;
  link?: string;
  icon?: string;
  router?: any;
}

const tmp_router = (title: string, icon: string, path: string, com:
  string) => {
  let name = 'app.' + path
  return {
    title: title,
    icon: icon,
    name: name,
    to: { name: name, params: { id: '' } },
    router: {
      path: path,
      name: name,
      component: () => import('../pages/app/' + com + '.vue')
    },
  }
}
let uniqueLinks: { [key: string]: [MenuLink] } = {
  [cfg.id]: [tmp_router('系统信息', 'v-data-view', 'oasys', 'oasys')]
}

const defaultLinks: MenuLink[] = [
  {
    title: '应用中心',
    caption: '',
    icon: 'v-apps',
    to: { name: 'home' }
  },
  {
    title: '文件管理',
    caption: '',
    icon: 'v-folder',
    to: { name: 'fs' }
  },
  {
    title: '账号设置',
    icon: 'v-user',
    to: { name: 'user' }
  },
  {
    title: '系统监控',
    icon: 'v-data-view',
    to: { name: 'stats' }
  },
  {
    title: '文档中心',
    icon: 'v-file-exception',
    to: { name: 'doc' }
  },
  {
    title: '设置',
    caption: '',
    icon: 'v-setting',
    to: { name: 'settings' }
  },
]

const items = ref(defaultLinks)

const load_default = () => {
  items.value = defaultLinks
}


const appLinks = ref([
  tmp_router('', 'v-home', 'home', 'home'),
  tmp_router('用户管理', 'v-team', 'user', 'user'),
  tmp_router('权限管理', 'v-key', 'auth', 'auth'),
  tmp_router('应用设置', 'v-setting', 'cfg', 'cfg'),
  tmp_router('test', 'v-key', 'test', '../IndexPage'),
] as MenuLink[])


export default { items, load_default, appLinks, uniqueLinks }

