/*
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:57
* @description：user
*/

import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
    state: () => {
        return {
            hideHeader: false,
            title: '',
            isDark: false,
        }
    },
    actions: {
    },
})
