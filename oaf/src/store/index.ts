import {InjectionKey} from 'vue'
import {createStore, useStore as baseUseStore, Store} from 'vuex'
import api from "../api";
import {User, UserState} from './user'

export interface State extends Object {
    oauuid: string
    user: UserState
    apps: []
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    modules: {
        user: User
    },
// @ts-ignore
    state: {
        oauuid: '',
        apps: []
    },
    getters: {},
    mutations: {
        setOA(state: any, data: any) {
            state.oauuid = data.uuid
        },
        setApps(state: State, data: any) {
            state.apps = data
        }
    },
    actions: {
        fetchSelf({commit}) {
            api.app.self().Start(d => {
                commit('setOA', d)
            })
        },
        fetchApps({commit}) {
            api.app.list().Start(e => {
                commit('setApps', e)
            })
        }
    }
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
    return baseUseStore(key)
}
