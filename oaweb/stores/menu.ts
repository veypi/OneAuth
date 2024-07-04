/*
 * menu.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-07 17:51
 * Distributed under terms of the MIT license.
 */

interface item {
  ico: string
  name: string
  path: string
  label?: string
  subs?: item[]
}

const default_menu = [
  { ico: 'home', name: '应用中心', path: '/' },
  { ico: 'user', name: '用户设置', path: '/user' },
  { ico: 'file-exception', name: '文档中心', path: '/docs' },
  { ico: 'data-view', name: '应用统计', path: '/stats' },
  { ico: 'setting', name: '系统设置', path: '/setting' },
]

export const useMenuStore = defineStore('menu', {
  state: () => ({
    menus: default_menu as item[]
  }),
  getters: {
  },
  actions: {
    set(list: item[]) {
      this.menus = list
    },
    default() {
      this.menus = default_menu
    }
  },
});

