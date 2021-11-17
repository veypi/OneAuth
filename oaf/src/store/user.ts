import {Module} from "vuex";
import api from "@/api";
import util from '@/libs/util'
import {Base64} from 'js-base64'
import {State} from './index'
import router from "@/router";
import {Auths, NewAuths} from '@/auth'

export interface UserState {
    id: number
    username: string
    nickname: string
    phone: string
    icon: string
    email: string
    ready: boolean
    auth: Auths

    [key: string]: any
}

export const User: Module<UserState, State> = {
    namespaced: true,
    state: {
        id: 0,
        username: '',
        nickname: '',
        phone: '',
        icon: '',
        email: '',
        auth: NewAuths([]),
        ready: false
    },
    mutations: {
        setBase(state: UserState, data: any) {
            state.id = data.id
            state.icon = data.icon
            state.username = data.username
            state.nickname = data.nickname
            state.phone = data.phone
            state.email = data.email
            state.ready = true
        },
        setAuth(state: UserState, data: any) {
            state.auth = NewAuths(data)
        },
        logout(state: UserState) {
            state.ready = false
            localStorage.removeItem('auth_token')
            router.push({name: 'login'})
        }
    },
    actions: {
        fetchUserData(context) {
            let token = util.getToken()?.split('.');
            if (!token || token.length !== 3) {
                return false
            }
            let data = JSON.parse(Base64.decode(token[1]))
            if (data.id > 0) {
                context.commit('setAuth', data.auth)
                api.user.get(data.id).Start(e => {
                    context.commit('setBase', e)
                })
            }
        }
    }
}
