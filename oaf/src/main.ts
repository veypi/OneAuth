import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import {Api} from '@/api'
import OneIcon from '@veypi/one-icon'
import Message from 'vue-m-message'
import 'vue-m-message/dist/index.css'

Vue.use(Message) // will mount `Vue.prototype.$message`

// Vue.use(OneIcon, {href: 'https://at.alicdn.com/t/font_2872366_7aws02sx9bl.js'})
Vue.use(OneIcon, {href: './icon.js'})
Vue.use(Api)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
