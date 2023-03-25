/*
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:57
* @description：user
*/

import { modelsBread } from '@/models'
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
    state: () => {
        return {
            id: '',
            hideHeader: false,
            title: '',
            isDark: false,
            breads: [] as modelsBread[],
        }
    },
    actions: {
        toggle_theme() {
            this.isDark = !this.isDark
            document.documentElement.setAttribute('theme', this.isDark ? 'dark' : '')
        },
        setBreads(b: modelsBread) {
            let l = this.breads.length
            for (let i = l; i < b.Index; i++) {
                this.breads.push({} as modelsBread)
            }
            this.breads[b.Index] = b
        },
    },
})
