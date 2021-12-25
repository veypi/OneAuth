import {Module} from "vuex";
import api from "@/api";
import util from '@/libs/util'
import {Base64} from 'js-base64'
import {State} from './index'
import router from "@/router";
import {Auths, NewAuths} from '@/auth'
import {modelsSimpleAuth, modelsUser} from '@/models'
import {Cfg} from '@/oaer'

export interface UserState {
    id: number
    local: modelsUser
    ready: boolean
    auth: Auths
}

export const User: Module<UserState, State> = {
    namespaced: true,
    state: {
        id: 0,
        local: {} as modelsUser,
        auth: NewAuths([]),
        ready: false
    },
    mutations: {
        setBase(state: UserState, data: modelsUser) {
            state.id = data.ID
            state.local = data
            state.ready = true
        },
        setAuth(state: UserState, data: modelsSimpleAuth[]) {
            state.auth = NewAuths(data)
        },
        refreshToken(state: UserState, data: string) {
            Cfg.token.value = util.getToken()
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
            if (data.ID > 0) {
                context.commit('setAuth', data.Auth)
                api.user.get(data.ID).Start(e => {
                    context.commit('setBase', e)
                },e=> {
                    context.commit('logout')
                })
            }
        }
    }
}
