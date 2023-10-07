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
