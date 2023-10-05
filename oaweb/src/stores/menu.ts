/*
 * menu.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 21:51
 * Distributed under terms of the MIT license.
 */


import { defineStore } from 'pinia';
import { MenuLink } from 'src/models';

const defaultLinks: MenuLink[] = [
  {
    title: '应用中心',
    caption: '',
    icon: 'apps',
    to: { name: 'home' }
  },
  {
    title: '用户管理',
    caption: 'oa.veypi.com',
    icon: 'person',
    to: { name: 'user' }
  },
  {
    title: '设置',
    caption: '',
    icon: 'settings',
    to: { name: 'settings' }
  },
]
export const useMenuStore = defineStore('menu', {
  state: () => ({
    list: defaultLinks as MenuLink[],
  }),
  getters: {
  },
  actions: {
    set(links: MenuLink[]) {
      this.list = links

    },
    load_default() {
      this.list = defaultLinks
    }
  },
});
