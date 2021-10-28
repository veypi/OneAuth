import {InjectionKey} from 'vue'
import {createStore, useStore as baseUseStore, Store} from 'vuex'
import api from "../api";
import router from "../router";
import {darkTheme} from 'naive-ui'

export interface State {
    oauuid: string
    user: object
    theme: string

}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state: {
        theme: 'light',
        oauuid: '',
        user: {}
    },
    getters: {
        IsDark(state: any) {
            return state.theme === 'dark'
        },
        GetTheme(state: any, getters) {
            return getters.IsDark ? darkTheme : null
        }
    },
    mutations: {
        setOA(state: any, data: any) {
            state.oauuid = data.uuid
        },
        setTheme(state: any, t: string) {
            state.theme = t
        }
    },
    actions: {
        changeTheme(context) {
            if (context.getters.IsDark) {
                context.commit('setTheme', 'light')
            } else {
                context.commit('setTheme', 'dark')
            }
        },
        fetchSelf({commit}) {
            api.app.self().Start(d => {
                commit('setOA', d)
            })
        },
        handleLogout() {
            localStorage.removeItem('auth_token')
            router.push({name: 'login'})
        }

    }
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
    return baseUseStore(key)
}
