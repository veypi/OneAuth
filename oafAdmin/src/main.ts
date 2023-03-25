import { createApp } from 'vue'
import App from './App.vue'
import OneIcon from '@veypi/one-icon'
import router from '@/router'
import { createPinia } from 'pinia'
import i18n from '@/i18n'
import 'animate.css'
import './assets/icon.js'
import './msg/index.css'
import './index.css'

let app = createApp(App)
app.use(OneIcon)
app.use(router)
app.use(i18n)
app.use(createPinia())
app.mount('#app')
