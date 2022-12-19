/*
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:57
* @description：user
*/

import {defineStore} from 'pinia'

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            id: 0,
            username: '',
        }
    },
    actions: {
        setUser() {
            this.id = 1
            this.username = 'admin'
        },
    },
})
