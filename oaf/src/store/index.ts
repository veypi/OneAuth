import {InjectionKey} from 'vue'
import {createStore, useStore as baseUseStore, Store} from 'vuex'
import api from '@/api'
import {User, UserState} from './user'
import {modelsBread} from '@/models'

type Map = { [key: string]: string }

export interface State extends Object {
    oauuid: string
    title: string
    user: UserState
    // 页面主页高度
    height: string
    hideHeader: boolean
    breads: modelsBread[]
    apps: []
    translateCache: Map
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    modules: {
        user: User,
    },
// @ts-ignore
    state: {
        oauuid: 'jU5Jo5hM',
        title: '',
        height: 'calc(100vh - 108px)',
        hideHeader: false,
        apps: [],
        translateCache: {},
        breads: [{Index: 0, Name: 'home', Type: 'icon', RName: 'home'}],
    },
    getters: {
        cache: (state: State) => (key: string) => {
            return state.translateCache[key] || key
        },
    },
    mutations: {
        setOA(state: any, data: any) {
            state.oauuid = data.uuid
        },
        setApps(state: State, data: any) {
            state.apps = data
        },
        addCache(state: State, data: Map) {
            Object.assign(state.translateCache, data)
        },
        hideHeader(state: State, i: boolean) {
            state.hideHeader = i
            if (i) {
                state.height = 'calc(100vh - 44px)'
            } else {
                state.height = 'calc(100vh - 108px)'
            }
        },
        setHeight(state: State, h: string) {
            state.height = h
        },
        setTitle(state: State, to: string) {
            state.title = to
        },
        setBreads(state: State, b: modelsBread) {
            let l = state.breads.length
            for (let i = l; i < b.Index; i++) {
                state.breads.push({} as modelsBread)
            }
            state.breads[b.Index] = b
        },
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
        },
    },
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
    return baseUseStore(key)
}
