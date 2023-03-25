/*
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2022-04-16 10:57
* @description：user
*/

import api from '@/api'
import { Auths, NewAuths } from '@/auth'
import { util } from '@/libs'
import { modelsUser } from '@/models'
import { Base64 } from 'js-base64'
import router from "@/router";
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            id: 0,
            auth: {} as Auths,
            local: {} as modelsUser,
            username: '',
        }
    },
    actions: {
        setUser() {
            this.id = 1
            this.username = 'admin'
        },
        fetchUserData() {
            let token = util.getToken()?.split('.');
            if (!token || token.length !== 3) {
                return false
            }
            let data = JSON.parse(Base64.decode(token[1]))
            console.log(data)
            if (data.id) {
                this.auth = NewAuths(data.access)
                console.log(this.auth)
                api.user.get(data.id).Start(e => {
                    console.log(e)
                    this.id = e.id
                }, e => {
                    this.logout()
                })
            }
        },
        logout() {
            localStorage.removeItem('auth_token')
            router.push({ name: 'login' })
        }
    },
})
