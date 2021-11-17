import Vue from 'vue'
import Vuex from 'vuex'
import api from '@/api'
import router from '@/router'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    oauuid: '',
    user: null
  },
  mutations: {
    setOA(state: any, data: any) {
      state.oauuid = data.uuid
    }
  },
  actions: {
    fetchSelf({commit}) {
      api.app.self().Start(d => {
        commit('setOA', d)
      })
    },
    handleLogout() {
      localStorage.removeItem('auth_token')
      router.push({name: 'login'})
    }
  },
  modules: {}
})
